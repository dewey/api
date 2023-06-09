package main

import "flag"

type flags struct {
   r_path string
   w_before string
   w_module string
}

func main() {
   var f flags
   flag.StringVar(&f.r_path, "r", "", "read path")
   flag.StringVar(&f.w_module, "m", "2a.pages.dev/tls", "write module")
   flag.StringVar(&f.w_before, "w", "", "write before")
   flag.Parse()
   if f.r_path != "" {
      err := f.read()
      if err != nil {
         panic(err)
      }
   } else if f.w_before != "" {
      err := f.write()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
