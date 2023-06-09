package main

import (
   "fmt"
   "os"
   "os/exec"
   "strings"
)

func (f flags) write() error {
   var after string
   i := strings.LastIndexAny(f.w_before, ".:")
   if i >= 0 {
      after = f.w_before[i+1:]
   } else {
      after = f.w_before
   }
   cmd := exec.Command(
      "gorename",
      "-from", fmt.Sprintf(`%q%v`, f.w_module, f.w_before),
      "-to", "_" + after,
   )
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
   fmt.Printf("%q\n", cmd.Args)
   return cmd.Run()
}
