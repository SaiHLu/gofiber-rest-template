package common

import "strconv"

func FormatValidationMessage(tag string, value interface{}) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}

	return ""
}

func ConvertStringToInt(val string) (int, error) {
	page, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return page, nil
}
