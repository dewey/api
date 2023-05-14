package main

import "os"

func do_gvimrc(home string) error {
   data, err := os.ReadFile("_gvimrc")
   if err != nil {
      return err
   }
   return os.WriteFile(home + "/_gvimrc", data, 0666)
}
