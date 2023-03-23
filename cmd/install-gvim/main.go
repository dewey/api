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
   var syntax bool
   flag.BoolVar(&syntax, "syntax", false, "install syntax")
   flag.Parse()
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   switch {
   case gvim:
      err = do_gvim(home)
   case gvimrc:
      err = do_gvimrc(home)
   case syntax:
      err = do_syntax()
   default:
      flag.Usage()
   }
   if err != nil {
      panic(err)
   }
}
