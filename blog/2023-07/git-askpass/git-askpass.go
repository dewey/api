package main

import (
   "encoding/json"
   "fmt"
   "os"
   "strings"
)

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   u, err := user(home + "/github.json")
   if err != nil {
      panic(err)
   }
   if len(os.Args) == 2 {
      prompt := os.Args[1]
      switch {
      case strings.HasPrefix(prompt, "Username"):
         fmt.Fprintln(os.Stderr, "Username")
         fmt.Println(u["username"])
      case strings.HasPrefix(prompt, "Password"):
         fmt.Fprintln(os.Stderr, "Password")
         fmt.Println(u["password"])
      }
   }
}
