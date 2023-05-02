package main

import (
   "fmt"
   "go/doc"
   "go/parser"
   "go/token"
)

func main() {
   fset := token.NewFileSet()
   pkgs, err := parser.ParseDir(fset, "D:/git/tls", nil, 0)
   if err != nil {
      panic(err)
   }
   pkg := doc.New(pkgs["tls"], "tls", 0)
   for _, typ := range pkg.Types {
      for _, method := range typ.Methods {
         fmt.Println(typ.Name, method.Name)
         {
            items := method.Decl.Type.Params
            if items != nil {
               for _, item := range items.List {
                  fmt.Printf("%#v\n", item.Type)
               }
            }
         }
         {
            items := method.Decl.Type.Results
            if items != nil {
               for _, item := range items.List {
                  fmt.Printf("%#v\n", item.Type)
               }
            }
         }
         fmt.Println()
      }
   }
}
