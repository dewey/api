package main

import (
   "encoding/json"
   "fmt"
   "os"
   "strings"
)

func credential(name string) (map[string]string, error) {
   data, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   var cred map[string]string
   if err := json.Unmarshal(data, &cred); err != nil {
      return nil, err
   }
   return cred, nil
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   cred, err := credential(home + "/Documents/github.json")
   if err != nil {
      panic(err)
   }
   if len(os.Args) == 2 {
      prompt := os.Args[1]
      switch {
      case strings.HasPrefix(prompt, "Username"):
         fmt.Fprintln(os.Stderr, "Username")
         fmt.Println(cred["username"])
      case strings.HasPrefix(prompt, "Password"):
         fmt.Fprintln(os.Stderr, "Password")
         fmt.Println(cred["password"])
      }
   }
}
