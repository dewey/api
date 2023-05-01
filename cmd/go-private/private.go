package main

import (
   "fmt"
   "os"
)

func main() {
   if len(os.Args) == 2 {
   } else {
      fmt.Println("private [identifier]")
   }
}
