package main

import (
   "fmt"
   "strconv"
   "time"
)

func main() {
   now := time.Now().Unix()
   id := strconv.FormatInt(now, 36)
   fmt.Println(id + "@riseup.net")
}
