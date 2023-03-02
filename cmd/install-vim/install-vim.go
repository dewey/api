package main

import (
   "fmt"
   "net/http"
   "os"
   "path/filepath"
)

var patches = []patch{
   // github.com/fleiner/vim/issues/2
   // github.com/vim/vim/pull/8023
   {"vim/vim/a942f9ad/runtime/", "syntax/javascript.vim"},
   // github.com/tpope/vim-markdown/pull/175
   {"tpope/vim-markdown/564d7436/", "syntax/markdown.vim"},
   // github.com/NLKNguyen/papercolor-theme/pull/167
   {"NLKNguyen/papercolor-theme/e397d18a/", "colors/PaperColor.vim"},
   // github.com/vim/vim/issues/11996
   {"google/vim-ft-go/master/", "syntax/go.vim"},
}

func download(in, out string) error {
   fmt.Println("GET", in)
   res, err := http.Get(in)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if err := os.MkdirAll(filepath.Dir(out), os.ModePerm); err != nil {
      return err
   }
   file, err := os.Create(out)
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}

const gvim =
   "https://github.com/vim/vim-win32-installer/releases/download/" +
   "v8.2.3526/gvim_8.2.3526_x64.zip"

type patch struct {
   dir, base string
}
