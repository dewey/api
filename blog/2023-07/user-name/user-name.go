package main

import (
   "bufio"
   "fmt"
   "golang.org/x/exp/slices"
   "os"
   "strings"
)

func main() {
   file, err := os.Open("american-english")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   buf := bufio.NewScanner(file)
   var words []string
   for buf.Scan() {
      before := buf.Text()
      if len(before) != 4 {
         continue
      }
      if strings.Contains(before, "'") {
         continue
      }
      after := buf.Bytes()
      slices.Sort(after)
      after = slices.Compact(after)
      if len(after) != 4 {
         continue
      }
      words = append(words, strings.ToLower(before))
   }
   slices.Sort(words)
   words = slices.Compact(words)
   for _, word := range words {
      fmt.Println(word)
   }
}
