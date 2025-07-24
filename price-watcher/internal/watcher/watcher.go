package watcher

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/maksroxx/DivineEye/price-watcher/internal/kafka"
	"github.com/maksroxx/DivineEye/price-watcher/internal/logger"
	"go.uber.org/zap"
)

type PriceWatcher interface {
	Run(ctx context.Context) error
}

type Watcher struct {
	Producer kafka.Producerer
	Symbols  []string
	Log      logger.Logger
}

func NewWatcher(producer kafka.Producerer, symbols []string, log logger.Logger) PriceWatcher {
	return &Watcher{
		Producer: producer,
		Symbols:  symbols,
		Log:      log,
	}
}

func chunkSymbols(symbols []string, chunkSize int) [][]string {
	var chunks [][]string
	for chunkSize < len(symbols) {
		symbols, chunks = symbols[chunkSize:], append(chunks, symbols[:chunkSize:chunkSize])
	}
	if len(symbols) > 0 {
		chunks = append(chunks, symbols)
	}
	return chunks
}

func (w *Watcher) Run(ctx context.Context) error {
	chunks := chunkSymbols(w.Symbols, 50)

	for _, symbolsChunk := range chunks {
		go func(chunk []string) {
			for {
				if err := w.runConnection(ctx, chunk); err != nil {
					w.Log.Error("[Watcher.go]", zap.String("chunk_failed", fmt.Sprintf("[reconnect] chunk failed, reconnecting... %v", err)))
					time.Sleep(5 * time.Second)
				}
			}
		}(symbolsChunk)
	}
	select {}
}

func (w *Watcher) runConnection(ctx context.Context, symbols []string) error {
	streams := []string{}
	for _, s := range symbols {
		streams = append(streams, fmt.Sprintf("%s@trade", strings.ToLower(s)))
	}
	streamQuery := strings.Join(streams, "/")

	u := url.URL{
		Scheme:   "wss",
		Host:     "stream.binance.com:9443",
		Path:     "/stream",
		RawQuery: "streams=" + streamQuery,
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	defer c.Close()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			_, msg, err := c.ReadMessage()
			if err != nil {
				return err
			}

			var parsed struct {
				Stream string `json:"stream"`
				Data   struct {
					Price string `json:"p"`
				} `json:"data"`
			}

			if err := json.Unmarshal(msg, &parsed); err != nil {
				continue
			}

			symbol := strings.Split(parsed.Stream, "@")[0]
			var price float64
			fmt.Sscanf(parsed.Data.Price, "%f", &price)

			w.Log.Info("[Watcher.go]", zap.String("symbol_price", fmt.Sprintf("[Binance] %s -> %.2f", symbol, price)))
			_ = w.Producer.PublishPrice(symbol, price)
		}
	}
}
