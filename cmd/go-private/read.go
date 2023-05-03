package main

import (
   "fmt"
   "go/ast"
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
                  ident := ast_ident(item.Type)
                  fmt.Println(is_exported(ident), ident)
               }
            }
         }
         {
            items := method.Decl.Type.Results
            if items != nil {
               for _, item := range items.List {
                  ident := ast_ident(item.Type)
                  fmt.Println(is_exported(ident), ident)
               }
            }
         }
         fmt.Println()
      }
   }
}

func is_exported(ident *ast.Ident) bool {
   if doc.IsPredeclared(ident.String()) {
      return true
   }
   return ident.IsExported()
}

func ast_ident(expr ast.Expr) *ast.Ident {
   switch v := expr.(type) {
   case *ast.ArrayType:
      return ast_ident(v.Elt)
   case *ast.Ident:
      return v
   case *ast.SelectorExpr:
      return v.Sel
   case *ast.StarExpr:
      return ast_ident(v.X)
   }
   panic(expr)
}
