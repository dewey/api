package main

import "os"

func do_config(home string) error {
   data, err := os.ReadFile(".gitconfig")
   if err != nil {
      return err
   }
   return os.WriteFile(home + "/.gitconfig", data, os.ModePerm)
}
