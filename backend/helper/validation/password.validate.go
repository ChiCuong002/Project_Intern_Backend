package validation

import "unicode"

func IsPassword(password string) (bool, map[string]string) {
	errors := make(map[string]string)
	hasNum, hasUpper, hasSpecial := false, false, false
	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			hasNum = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	if len(password) < 8 {
		errors["length"] = "Password length should be at least 8 characters"
	}
	if !hasNum {
		errors["num"] = "Password should have at least one number"
	}
	if !hasUpper {
		errors["upper"] = "Password should have at least one uppercase letter"
	}
	if !hasSpecial {
		errors["special"] = "Password should have at least one special character"
	}
	return len(errors) == 0, errors
}
