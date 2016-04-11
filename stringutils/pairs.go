package stringutils

import (
  "fmt"
)

func Pairs(s []string) (pairs []string) {
  for index, value := range s {
    if index == 0 {
      continue
    }
    
    pair := fmt.Sprintf("%s%s", s[index-1], value)
    pairs = append(pairs, pair)
  }
  return pairs
}