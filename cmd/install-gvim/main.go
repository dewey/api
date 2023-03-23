package main

import (
   "flag"
   "os"
)

func main() {
   var gvim bool
   flag.BoolVar(&gvim, "gvim", false, "install GVIM")
   var gvimrc bool
   flag.BoolVar(&gvimrc, "gvimrc", false, "install GVIMRC")
   flag.Parse()
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   if gvimrc {
      err := do_gvimrc(home)
      if err != nil {
         panic(err)
      }
   } else if gvim {
      err := do_gvim(home)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
