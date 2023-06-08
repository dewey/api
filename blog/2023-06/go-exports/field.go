package main

import "go/ast"

func field_switch(f *ast.Field) {
   for _, name := range f.Names {
      name.Name = to_lower(name.Name)
   }
   switch a := f.Type.(type) {
   case *ast.ArrayType:
   case *ast.FuncType:
      for _, field := range a.Params.List {
         field_switch(field)
      }
      if a.Results != nil {
         for _, field := range a.Results.List {
            field_switch(field)
         }
      }
   case *ast.Ident:
      a.Name = to_lower(a.Name)
   case *ast.SelectorExpr:
   case *ast.StarExpr:
      switch b := a.X.(type) {
      case *ast.Ident:
         b.Name = to_lower(b.Name)
      default:
         panic(b)
      }
   default:
      panic(a)
   }
}
