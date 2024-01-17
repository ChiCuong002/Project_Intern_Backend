package validation

import "regexp"

func IsPhoneNumber(phonenumber string) bool {
	re := regexp.MustCompile(`^0[0-9]{9}$`)
	return re.MatchString(phonenumber)
}
