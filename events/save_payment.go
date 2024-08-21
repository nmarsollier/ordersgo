package events

// SavePayment saves a payment event
func SavePayment(data *PaymentEvent, ctx ...interface{}) (*Event, error) {
	event, err := insert(newPaymentEvent(data), ctx...)

	if err != nil {
		return nil, err
	}

	return event, nil
}
