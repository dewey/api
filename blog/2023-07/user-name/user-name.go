package main

import (
   "bufio"
   "fmt"
   "golang.org/x/exp/slices"
   "os"
   "strings"
)

func main() {
   file, err := os.Open("american-english-insane")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   buf := bufio.NewScanner(file)
   var words []string
   for buf.Scan() {
      word := strings.ToLower(buf.Text())
      if len(word) != 3 {
         continue
      }
      if word[0] == word[1] {
         continue
      }
      if word[0] == word[2] {
         continue
      }
      if word[1] == word[2] {
         continue
      }
      if strings.Contains(word, "'") {
         continue
      }
      if strings.Contains(word, "l") {
         continue
      }
      words = append(words, word)
   }
   slices.Sort(words)
   slices.Compact(words)
   for _, word := range words {
      fmt.Println(word)
   }
}
