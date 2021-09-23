package common

import "regexp"

var IsAlphabets = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
var IsAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
var IsNumeric = regexp.MustCompile(`^[0-9]+$`).MatchString
