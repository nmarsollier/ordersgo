package events

// SaveArticleExist saves the event for article exist
func SaveArticleExist(data *ValidationEvent, ctx ...interface{}) (*Event, error) {
	event, err := insert(newValidationEvent(data), ctx...)

	if err != nil {
		return nil, err
	}

	return event, nil
}
