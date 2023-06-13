package godocs

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   "2a.pages.dev/rosso",
   "2a.pages.dev/rosso/dash",
   "github.com/n7olkachev/imgdiff",
   "github.com/n7olkachev/imgdiff/pkg/imgdiff",
}

func Test_GoDocs(t *testing.T) {
   for _, test := range tests {
      rep, err := repository(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", rep)
      time.Sleep(time.Second)
   }
}
