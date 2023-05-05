package gerr

/**
 * Translate
 */
func T(errorCode uint32) string {
	switch errorCode {
	//////////////////////////
	// Client-side
	//////////////////////////
	case ErrorBindData:
		return "Failed to bind data"
	case ErrorValidData:
		return "Failed to valid data"
	//////////////////////////
	// Server-side
	//////////////////////////
	case ErrorConnect:
		return "Failed to connect database"
	case ErrorSaveData:
		return "Failed to save data"
	case ErrorRetrieveData:
		return "Failed to retrieve data"
	case ErrorLogin:
		return "Failed to login. Please try again!"
	case ErrorGetUserInfo:
		return "Failed to get userInfo. Please try again!"
	}

	return "Unknown error"
}
