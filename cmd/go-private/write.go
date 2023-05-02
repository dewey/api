package main

import (
   "fmt"
   "os"
   "os/exec"
   "strings"
)

func main() {
   if len(os.Args) == 2 {
      before := os.Args[1]
      after := strings.ToLower(before[:1]) + before[1:]
      cmd := exec.Command("gofmt", "-w", "-r", before+" -> "+after, ".")
      cmd.Stdout = os.Stdout
      cmd.Stderr = os.Stderr
      fmt.Printf("%q\n", cmd.Args)
      err := cmd.Run()
      if err != nil {
         panic(err)
      }
   } else {
      fmt.Println("go-private [identifier]")
   }
}
