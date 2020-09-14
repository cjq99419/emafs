package tools

func ConvertToInt(i interface{}) int {
	if i == nil {
		return -12345
	} else {
		return int(i.(float64))
	}
}

func ConvertToString(i interface{}) string {
	if i == nil {
		return ""
	} else {
		return i.(string)
	}
}
