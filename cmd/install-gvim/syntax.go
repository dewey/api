package main

import "os"

func do_syntax() error {
   in, err := os.Open("syntax/typst.vim")
   if err != nil {
      return err
   }
   defer in.Close()
   out, err := os.Create("D:/vim/vimfiles/syntax/typst.vim")
   if err != nil {
      return err
   }
   defer out.Close()
   if _, err := out.ReadFrom(in); err != nil {
      return err
   }
   return nil
}
