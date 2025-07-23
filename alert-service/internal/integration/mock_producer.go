package integration

type MockProducer struct {
	called bool
}

func (m *MockProducer) PublishAlertCreated(alertId, userId, coin, direction string, price float64) error {
	m.called = true
	return nil
}

func (m *MockProducer) PublishAlertDeleted(alertId, userId string) error {
	m.called = true
	return nil
}

func (m *MockProducer) Close() error { return nil }
