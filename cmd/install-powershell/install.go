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
      data, err := os.ReadFile(filepath.Base(name))
      if err != nil {
         panic(err)
      }
      if err := os.WriteFile(name, data, 0777); err != nil {
         panic(err)
      }
   }
}
