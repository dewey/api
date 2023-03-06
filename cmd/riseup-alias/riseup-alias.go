package main

import (
   "fmt"
   "strconv"
   "time"
)

func format_int(f float64) string {
   return strconv.FormatInt(int64(f), 36)
}

func main() {
   now := time.Now()
   then := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
   for {
      dur := now.Sub(then)
      id := format_int(dur.Hours())
      if len(id) >= 4 {
         break
      }
      fmt.Println(id)
      id = format_int(dur.Minutes())
      fmt.Println(id)
      id = format_int(dur.Seconds())
      fmt.Print(id, "\n\n")
      then = then.AddDate(-1, 0, 0)
   }
}
