package main

import (
   "2a.pages.dev/nursery/exports"
   "fmt"
   "os"
   "sort"
)

func main() {
   if len(os.Args) == 2 {
      exps, err := exports.Exports(os.Args[1])
      if err != nil {
         panic(err)
      }
      sort.Slice(exps, func(i, j int) bool {
         return fmt.Sprintf("%p", exps[i]) < fmt.Sprintf("%p", exps[j])
      })
      for _, exp := range exps {
         fmt.Printf("rf 'mv %v %u'\n", exp, exp)
      }
   } else {
      fmt.Println("exports [directory]")
   }
}
