package main

import (
   "2a.pages.dev/nursery/justwatch"
   "flag"
   "time"
)

func main() {
   var f flags
   flag.StringVar(&f.address, "a", "", "address")
   flag.StringVar(&f.language, "lang", "en", "language")
   flag.DurationVar(&f.sleep, "s", 99*time.Millisecond, "sleep")
   flag.BoolVar(&f.verbose, "v", false, "verbose")
   flag.Parse()
   if f.verbose {
      justwatch.Client.Log_Level = 2
   } else {
      justwatch.Client.Log_Level = 0
   }
   if f.address != "" {
      err := f.stream()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
