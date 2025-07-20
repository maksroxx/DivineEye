package integration

type MockProducer struct {
	called bool
}

func (m *MockProducer) PublishAlertCreated(alertId, userId string) error {
	m.called = true
	return nil
}

func (m *MockProducer) Close() error { return nil }
