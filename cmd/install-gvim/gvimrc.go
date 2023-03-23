package main

import "os"

func do_gvimrc(home string) error {
   in, err := os.Open("_gvimrc")
   if err != nil {
      return err
   }
   defer in.Close()
   out, err := os.Create(home + "/_gvimrc")
   if err != nil {
      return err
   }
   defer out.Close()
   if _, err := out.ReadFrom(in); err != nil {
      return err
   }
   return nil
}
