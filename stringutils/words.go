package stringutils

import (
  "strings"
  "regexp"
)

func Words(s string) (words []string) {
  f := func(c rune) bool {
    re, _ := regexp.Compile(`\W`)
    return re.FindString(string(c)) != ""
  }

  for _, word := range strings.FieldsFunc(s, f) {
    words = append(words, word)
  }
  return words
}