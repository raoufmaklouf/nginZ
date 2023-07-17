package main

import (
	"regexp"
	"strings"
)

func fuzzUrl(url string) string {
	fuzzurl := strings.Split(Rev(url), "/")[0]
	finalUrl := Rev(strings.ReplaceAll(Rev(url), fuzzurl, fuzzurl+"laertoNzzuf"))

	return finalUrl
}

func Rev(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isUrl(url string) bool {
	s := false

	if strings.HasPrefix(url, "http://") == true || strings.HasPrefix(url, "https://") == true {
		s = true
	}
	return s
}

func xMatch(rg string, str string) bool {
	match, _ := regexp.MatchString(rg, str)
	if match == true {
		return true
	} else {
		return false

	}

}
