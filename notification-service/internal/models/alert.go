package models

type Alert struct {
	ID        string
	UserID    string
	Coin      string
	Price     float64
	Direction string
}
