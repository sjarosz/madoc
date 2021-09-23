package common

import "regexp"

var (
	AlphaRegex = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	NumRegex   = regexp.MustCompile(`^[0-9]+$`).MatchString
)
