package events

// SaveArticleExist saves the event for article exist
func SaveArticleExist(data *ValidationEvent, deps ...interface{}) (*Event, error) {
	event, err := insert(newValidationEvent(data), deps...)

	if err != nil {
		return nil, err
	}

	return event, nil
}
