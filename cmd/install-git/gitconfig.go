package main

import "os"

func do_config(home string) error {
   in, err := os.Open(".gitconfig")
   if err != nil {
      return err
   }
   defer in.Close()
   out, err := os.Create(home + "/.gitconfig")
   if err != nil {
      return err
   }
   defer out.Close()
   if _, err := out.ReadFrom(in); err != nil {
      return err
   }
   return nil
}
