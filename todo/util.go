package todo

// isValid return appropriate message if the given todo struct is invalid, else returns blank string
func isValid(t CreateTodo) (invalidMsg string) {

	// blank fields check
	if t.Title == "" || t.Status == "" {
		return "Todo request is missing status or title"
	}

	// valid status check
	statusValid := false
	for _, s := range allowedStatuses {
		if t.Status == s {
			statusValid = true
			break
		}
	}
	if !statusValid {
		return "The provided status is not supported"
	}

	//
	return ""
}
