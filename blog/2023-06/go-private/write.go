package main

import (
   "fmt"
   "os"
   "os/exec"
)

func run(name string, arg ...string) error {
   cmd := exec.Command(name, arg...)
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
   fmt.Printf("%q\n", cmd.Args)
   return cmd.Run()
}

func (f flags) write() error {
   err := run("gofmt", "-w", "-r", f.w_before+" -> _"+f.w_before, ".")
   if err != nil {
      return err
   }
   return run("gofmt", "-w", "-r", "net._"+f.w_before+" -> net."+f.w_before, ".")
}
