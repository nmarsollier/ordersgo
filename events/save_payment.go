package events

// SavePayment saves a payment event
func SavePayment(data *PaymentEvent) (*Event, error) {
	event, err := insert(newPaymentEvent(data))

	if err != nil {
		return nil, err
	}

	return event, nil
}
