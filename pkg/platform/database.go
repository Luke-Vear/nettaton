package platform

// UpdateUserScore will increase the users overall and questions specific score by 1.
func UpdateUserScore(questionType string, userID string, correct bool) error {

	if userID == "" {
		return nil
	}

	return nil
}
