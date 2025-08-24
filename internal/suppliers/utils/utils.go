package utils

import "strings"

type Suppliers string

const (
	Acme       Suppliers = "acme"
	Patagonia  Suppliers = "patagonia"
	Paperflies Suppliers = "paperflies"
)

func TrimSpacesInString(s string) string {
	return strings.Trim(s, " ")
}

func TrimSpacesInSlices(s []string) []string {
	var result []string
	for _, str := range s {
		trimmed := TrimSpacesInString(str)
		result = append(result, trimmed)
	}
	return result
}
