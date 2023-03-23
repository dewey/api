package main

import (
   "2a.pages.dev/nursery"
   "fmt"
   "os"
   "path/filepath"
)

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.Join(home, "nursery/vim")
   home = filepath.Join(home, filepath.Base(gvim))
   fmt.Println("Stat", home)
   if _, err := os.Stat(home); err != nil {
      err := download(gvim, home)
      if err != nil {
         panic(err)
      }
   }
   fmt.Println("Zip", home)
   if err := nursery.Zip(home, `D:\vim`, 2); err != nil {
      panic(err)
   }
   for _, pat := range patches {
      err := download(
         "https://raw.githubusercontent.com/" + pat.dir + pat.base,
         filepath.Join(`D:\vim`, pat.base),
      )
      if err != nil {
         panic(err)
      }
   }
}
