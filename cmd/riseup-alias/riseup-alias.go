package main

import (
   "flag"
   "fmt"
   "strconv"
   "time"
)

func main() {
   var year int
   flag.IntVar(&year, "y", 0, "year")
   flag.Parse()
   if year >= 1 {
      then := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
      { // seconds
         dur := time.Now().Sub(then).Seconds()
         id := strconv.FormatInt(int64(dur), 36)
         fmt.Println(id + "@riseup.net")
      }
      { // minutes
         dur := time.Now().Sub(then).Minutes()
         id := strconv.FormatInt(int64(dur), 36)
         fmt.Println(id + "@riseup.net")
      }
      { // hours
         dur := time.Now().Sub(then).Hours()
         id := strconv.FormatInt(int64(dur), 36)
         fmt.Println(id + "@riseup.net")
      }
   } else {
      flag.Usage()
   }
}
