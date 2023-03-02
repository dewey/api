package nursery

import (
   "fmt"
   "testing"
)

type test_type struct {
   s string
   n int
}

var tests = []test_type{
   {"hello world", 0},
   {"/one/two", 2},
}

func Test_Strip(t *testing.T) {
   for _, test := range tests {
      s := strip(test.s, test.n)
      fmt.Println(s)
   }
}
