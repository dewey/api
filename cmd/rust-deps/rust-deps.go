package main

import (
   "bufio"
   "flag"
   "fmt"
   "os"
   "os/exec"
   "strings"
)

type flags struct {
   crate string
   features string
   tree bool
}

func (f flags) cargo_add() error {
   arg := []string{"add", "--no-default-features"}
   if f.features != "" {
      arg = append(arg, "--features", f.features)
   }
   arg = append(arg, f.crate)
   cmd := exec.Command("cargo", arg...)
   cmd.Stderr = os.Stderr
   fmt.Println(cmd)
   return cmd.Run()
}

func (f flags) cargo_tree() error {
   tree := exec.Command("cargo", "tree", "--prefix", "none")
   pipe, err := tree.StdoutPipe()
   if err != nil {
      return err
   }
   fmt.Println(tree)
   tree.Start()
   defer tree.Wait()
   scan := bufio.NewScanner(pipe)
   var count int
   deps := make(set)
   for scan.Scan() {
      text := scan.Text()
      dep, _, ok := strings.Cut(text, " ")
      if ok && deps.add(dep) {
         fmt.Print(count)
         count++
      }
      fmt.Print("\t", text, "\n")
   }
   return nil
}

type set map[string]struct{}

func (s set) add(t string) bool {
   _, ok := s[t]
   if ok {
      return false
   }
   s[t] = struct{}{}
   return true
}

func main() {
   var f flags
   flag.StringVar(&f.crate, "c", "", "crate")
   flag.StringVar(&f.features, "f", "", "features")
   flag.BoolVar(&f.tree, "t", false, "tree")
   flag.Parse()
   if f.crate != "" {
      err := f.cargo_add()
      if err != nil {
         panic(err)
      }
   } else if f.tree {
      err := f.cargo_tree()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
