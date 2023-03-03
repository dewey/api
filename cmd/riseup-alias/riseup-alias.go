package main

import (
   "fmt"
   "strconv"
   "time"
)

func main() {
   // wikipedia.org/wiki/GitHub
   then := time.Date(2008, 2, 8, 0, 0, 0, 0, time.UTC)
   dur := time.Now().Sub(then).Seconds()
   id := strconv.FormatInt(int64(dur), 36)
   fmt.Println(id + "@riseup.net")
}
