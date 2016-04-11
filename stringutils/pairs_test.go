package stringutils_test

import (
  "github.com/ivanfoong/stringutils-go/stringutils"
  "reflect"
  "testing"
)

func TestPairs(t *testing.T) {
  var testCases = []struct {
    input []string
    output []string
  }{
    {[]string{"a", "b", "c", "d", "e"}, []string{"ab", "bc", "cd", "de"}},
  }
  
  for _, tt := range testCases {
    input := tt.input
    expectedOutput := tt.output
    output := stringutils.Pairs(input)
    if !reflect.DeepEqual(expectedOutput, output) {
      t.Errorf("expected Pairs(%s) to equal `%v`, got `%v` instead", input, expectedOutput, output)
    } 
  }
}