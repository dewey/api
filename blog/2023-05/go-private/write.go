package main

import (
   "fmt"
   "os"
   "os/exec"
   "strings"
)

func (f flags) write() error {
   after := strings.ToLower(f.w_before[:1]) + f.w_before[1:]
   cmd := exec.Command("gofmt", "-w", "-r", f.w_before+" -> "+after, ".")
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
   fmt.Printf("%q\n", cmd.Args)
   return cmd.Run()
}
