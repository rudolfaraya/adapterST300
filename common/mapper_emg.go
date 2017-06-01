package common

func MapperEmg(emgID int64) string {

	switch emgID {
	case 1:
		message := "Emergency by panic button"
		return message
	case 2:
		message := "Emergency by parking lock"
		return message
	case 3:
		message := "Emergency by removing the main power"
		return message
	case 5:
		message := "Emergency by anti-theft"
		return message
	case 6:
		message := "Emergency by anti-theft door"
		return message
	case 7:
		message := "Emergency by motion"
		return message
	case 8:
		message := "Emergency by anti-theft shock"
		return message
	default:
		message := "NO ALERT"
		return message
	}
}
