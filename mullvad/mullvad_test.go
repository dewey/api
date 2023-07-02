package mullvad

import (
   "fmt"
   "os"
   "testing"
)

func Test_Server(t *testing.T) {
   file, err := os.Open("__data.json")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   servs, err := servers(file)
   if err != nil {
      t.Fatal(err)
   }
   for _, serv := range servs {
      fmt.Print(serv, "\n\n")
   }
}
