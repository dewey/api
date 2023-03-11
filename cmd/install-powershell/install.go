package main

import (
   "os"
   "path/filepath"
)

var names = []string{
   `C:\Users\Steven\AppData\Local\Microsoft\Windows Terminal\settings.json`,
   `C:\Users\Steven\Documents\PowerShell\Microsoft.PowerShell_profile.ps1`,
}

func main() {
   for _, name := range names {
      in, err := os.Open(filepath.Base(name))
      if err != nil {
         panic(err)
      }
      out, err := os.Create(name)
      if err != nil {
         panic(err)
      }
      if _, err := out.ReadFrom(in); err != nil {
         panic(err)
      }
      if err := in.Close(); err != nil {
         panic(err)
      }
      if err := out.Close(); err != nil {
         panic(err)
      }
   }
}
