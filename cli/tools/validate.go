package tools

import (
	"strconv"
	"strings"
)

func validatePass(p string) (bool, string) {
	if len(p) > 0 {
		return true, ""
	} else {
		return false, "Invalid password length"
	}
}

func validateUsername(u string) (bool, string) {
	if len(u) > 0 {
		return true, ""
	} else {
		return false, "Invalid username length"
	}
}

func validateEmail(e string) (bool, string) {
	aite := strings.Index(e, "@")
	point := strings.Index(e, ".")
	if len(e) > 0 && aite != -1 && point != -1 && point > aite {
		return true, ""
	} else {
		return false, "Invalid email format"
	}
}

func validateId(i string) (bool, string) {
	for i1 := range i {
		if !(i1 < '9' && i1 > '0') {
			return false, "Invalid ID"
		}
	}
	return true, ""
}

func validatePhone(p string) (bool, string) {
	if len(p) != 11 {
		return false, "Invalid phone length"
	}
	for i1 := range p {
		if !(i1 < '9' && i1 > '0') {
			return false, "Invalid phone number"
		}
	}
	return true, ""
}

func validateLimit(n string) (bool, string) {
	if len(n) < 0 {
		return false, "Empty limit"
	}
	if i, err := strconv.Atoi(n); err != nil || i <= 0 {
		return false, "Invalid limit"
	}
	return true, ""
}