package general

import (
	"regexp"
	"strings"
)

func IsEmailValid(email string) bool {
	if email == "" {
		return false
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func IsPasswordStrong(password string) bool {
	if len(password) < 6 {
		return false
	}
	return true
}

func IsUsernameValid(username string) bool {
	if len(username) < 3 {
		return false
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return re.MatchString(username)
}

func NormalizeString(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}
