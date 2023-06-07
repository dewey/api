package main

import (
   "fmt"
   "go/ast"
   "go/parser"
   "go/printer"
   "go/token"
   "os"
   "strings"
)

func to_lower(s string) string {
   if s == "" {
      return ""
   }
   if len(s) == 1 {
      return strings.ToLower(s)
   }
   return strings.ToLower(s[:1]) + s[1:]
}

func main() {
   fset := token.NewFileSet()
   file, err := parser.ParseFile(fset, "tls/tls.go", nil, 0)
   if err != nil {
      panic(err)
   }
   for _, decl := range file.Decls {
      switch a := decl.(type) {
      case *ast.FuncDecl:
         a.Name.Name = to_lower(a.Name.Name)
         for _, field := range a.Type.Params.List {
            field_switch(field)
         }
         for _, field := range a.Type.Results.List {
            field_switch(field)
         }
         printer.Fprint(os.Stdout, fset, a)
         fmt.Print("\n\n")
      case *ast.GenDecl:
         for _, spec := range a.Specs {
            switch b := spec.(type) {
            case *ast.ImportSpec:
            case *ast.TypeSpec:
               switch c := b.Type.(type) {
               case *ast.Ident:
               case *ast.InterfaceType:
                  for _, field := range c.Methods.List {
                     field_switch(field)
                  }
               case *ast.StructType:
                  for _, field := range c.Fields.List {
                     field_switch(field)
                  }
               default:
                  panic(c)
               }
               b.Name.Name = to_lower(b.Name.Name)
            case *ast.ValueSpec:
               for _, name := range b.Names {
                  name.Name = to_lower(name.Name)
               }
            default:
               panic(b)
            }
         }
         printer.Fprint(os.Stdout, fset, a)
         fmt.Print("\n\n")
      default:
         panic(a)
      }
   }
}
