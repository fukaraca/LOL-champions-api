package lib

import "strings"

//Striper is a function for trimming whitespaces till input becomes nil if it's necessary
func Striper(str string) *string {
	str = strings.TrimSpace(str)
	if str == "" {
		return nil
	}
	return &str
}
