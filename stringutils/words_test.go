package stringutils_test

import (
  "github.com/ivanfoong/stringutils-go/stringutils"
  "reflect"
  "testing"
)

func TestWords(t *testing.T) {
  var testCases = []struct {
    input string
    output []string
  }{
    {"testing 1 2 3", []string{"testing", "1", "2", "3"}},
    {"testing 1,2,3", []string{"testing", "1", "2", "3"}},
  }
  
  for _, tt := range testCases {
    input := tt.input
    expectedOutput := tt.output
    output := stringutils.Words(input)
    if !reflect.DeepEqual(expectedOutput, output) {
      t.Errorf("expected Words(%s) to equal `%v`, got `%v` instead", input, expectedOutput, output)
    } 
  }
}
