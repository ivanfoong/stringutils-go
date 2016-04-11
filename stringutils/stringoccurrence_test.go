package stringutils_test

import (
  "github.com/ivanfoong/stringutils-go/stringutils"
  "reflect"
  "testing"
)

func TestStringOccurrence(t *testing.T) {
  var testCases = []struct {
    input []string
    output map[string]int
  }{
    {[]string{"word", "word", "word1", "testing", "testing3"}, map[string]int{"word": 2, "word1": 1, "testing": 1, "testing3": 1}},
  }
  
  for _, tt := range testCases {
    input := tt.input
    expectedOutput := tt.output
    output := stringutils.StringOccurrence(input)
    if !reflect.DeepEqual(expectedOutput, output) {
      t.Errorf("expected StringOccurrence(%s) to equal `%v`, got `%v` instead", input, expectedOutput, output)
    } 
  }
}