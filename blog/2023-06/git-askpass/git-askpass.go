package main

import (
   "fmt"
   "os"
   "strings"
)

func user_info(name string) ([]string, error) {
   text, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   return strings.Split(string(text), "\n"), nil
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   user, err := user_info(home + "/2a/nursery/github.txt")
   if err != nil {
      panic(err)
   }
   if len(os.Args) == 2 {
      prompt := os.Args[1]
      switch {
      case strings.HasPrefix(prompt, "Username"):
         fmt.Fprintln(os.Stderr, "Username")
         fmt.Println(user[0])
      case strings.HasPrefix(prompt, "Password"):
         fmt.Fprintln(os.Stderr, "Password")
         fmt.Println(user[1])
      }
   }
}
