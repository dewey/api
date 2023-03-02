package main

import (
   "os"
   "path/filepath"
)

func main() {
   in, err := os.Open("Microsoft.PowerShell_profile.ps1")
   if err != nil {
      panic(err)
   }
   defer in.Close()
   name := in.Name()
   name = filepath.Join(`C:\Users\Steven\Documents\PowerShell`, name)
   out, err := os.Create(name)
   if err != nil {
      panic(err)
   }
   defer out.Close()
   if _, err := out.ReadFrom(in); err != nil {
      panic(err)
   }
}
