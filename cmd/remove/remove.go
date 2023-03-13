package main

import (
   "fmt"
   "os"
)

var names = []string{
   `C:\Users\Steven\.android`,
   `C:\Users\Steven\.cargo`,
   `C:\Users\Steven\AppData\Roaming\Mozilla\Firefox\Profiles\7dnqqks1.default-release\storage\default`,
   `C:\Users\Steven\go\pkg`,
}

func main() {
   for _, name := range names {
      _, err := os.Stat(name)
      if err != nil {
         fmt.Println(err)
      } else {
         err := os.RemoveAll(name)
         if err != nil {
            panic(err)
         }
      }
   }
}
