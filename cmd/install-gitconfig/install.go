package main

import (
   "os"
)

func main() {
   in, err := os.Open(".gitconfig")
   if err != nil {
      panic(err)
   }
   defer in.Close()
   out, err := os.Create(`C:\Users\Steven\.gitconfig`)
   if err != nil {
      panic(err)
   }
   defer out.Close()
   if _, err := out.ReadFrom(in); err != nil {
      panic(err)
   }
}
