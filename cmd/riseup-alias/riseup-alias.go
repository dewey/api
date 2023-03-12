package main

import (
   "fmt"
   "sort"
   "time"
)

func main() {
   now := time.Now()
   then := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
   var names []string
   for y := 0; y < 9; y++ {
      dur := now.Sub(then.AddDate(-y, 0, 0))
      names = append(names, hours(dur))
      names = append(names, minutes(dur))
      names = append(names, seconds(dur))
   }
   sort.Slice(names, func(i, j int) bool {
      return len(names[i]) < len(names[j])
   })
   for _, name := range names {
      fmt.Println(name)
   }
}

// Username may only contain alphanumeric characters or single hyphens, and
// cannot begin or end with a hyphen.
const digits = "23456789abcdefghijkmnpqrstuvwxyz"

func format_bits(u uint64) string {
   var a [64]byte
   i := len(a)
   b := uint64(len(digits))
   for u >= b {
      i--
      q := u / b
      a[i] = digits[u-q*b]
      u = q
   }
   i--
   a[i] = digits[u]
   return string(a[i:])
}

func hours(d time.Duration) string {
   f := d.Hours()
   return format_bits(uint64(f))
}

func minutes(d time.Duration) string {
   f := d.Minutes()
   return format_bits(uint64(f))
}

func seconds(d time.Duration) string {
   f := d.Seconds()
   return format_bits(uint64(f))
}
