package events

// SavePayment saves a payment event
func SavePayment(data *PaymentEvent, deps ...interface{}) (*Event, error) {
	event, err := insert(newPaymentEvent(data), deps...)

	if err != nil {
		return nil, err
	}

	return event, nil
}
