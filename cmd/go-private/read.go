package main

import (
   "fmt"
   "go/ast"
   "go/doc"
   "go/parser"
   "go/token"
)

func (f flags) read() error {
   fset := token.NewFileSet()
   pkgs, err := parser.ParseDir(fset, f.r_path, nil, 0)
   if err != nil {
      return err
   }
   pkg := doc.New(pkgs["tls"], "tls", 0)
   for _, fun := range pkg.Funcs {
      if !is_exported(fun.Decl.Type) {
         fmt.Println(fun.Name)
      }
   }
   for _, typ := range pkg.Types {
      for _, method := range typ.Methods {
         if !is_exported(method.Decl.Type) {
            fmt.Println(typ.Name, method.Name)
         }
      }
   }
   return nil
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
   return nil
}

func is_exported(ft *ast.FuncType) bool {
   iterate := func(items *ast.FieldList) bool {
      if items != nil {
         for _, item := range items.List {
            ident := ast_ident(item.Type)
            if !ident.IsExported() {
               if !doc.IsPredeclared(ident.String()) {
                  return false
               }
            }
         }
      }
      return true
   }
   if iterate(ft.Params) {
      if iterate(ft.Results) {
         return true
      }
   }
   return false
}
