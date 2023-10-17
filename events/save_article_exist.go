package events

// SaveArticleExist saves the event for article exist
func SaveArticleExist(data *ValidationEvent) (*Event, error) {
	event, err := insert(newValidationEvent(data))

	if err != nil {
		return nil, err
	}

	return event, nil
}
