/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package util

import (
	"regexp"
	"bytes"
)

var (
	wordRegex = regexp.MustCompile("[0-9A-Za-z]+")
)

func UppercaseFirstLetter(src string) string {
	words := wordRegex.FindAll([]byte(src),-1)
	for i, val := range words {
		words[i] = bytes.Title(val)
	}
	return string(bytes.Join(words,nil))
}
