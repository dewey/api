package main

import (
   "flag"
   "fmt"
   "time"
)

func main() {
   var length int
   flag.IntVar(&length, "len", 0, "length")
   var step int64
   flag.Int64Var(&step, "step", 999, "step")
   flag.Parse()
   if length >= 1 {
      now := time.Now()
      offset := now.Sub(now.Truncate(time.Second)).Milliseconds()
      for {
         s := format_bits(offset)
         if len(s) == length {
            fmt.Println(s)
         } else if len(s) > length {
            break
         }
         offset += step
      }
   } else {
      flag.Usage()
   }
}

// Username may only contain alphanumeric characters or single hyphens, and
// cannot begin or end with a hyphen.
const digits = "23456789abcdefghijkmnpqrstuvwxyz"

func format_bits(u int64) string {
   var a [64]byte
   i := len(a)
   b := int64(len(digits))
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
