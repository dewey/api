package main

import (
   "bufio"
   "fmt"
   "net/url"
   "os"
   "path/filepath"
   "strings"
)

func credentials(name string) ([]url.URL, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   var refs []url.URL
   buf := bufio.NewScanner(file)
   for buf.Scan() {
      var ref url.URL
      err := ref.UnmarshalBinary(buf.Bytes())
      if err != nil {
         return nil, err
      }
      refs = append(refs, ref)
   }
   return refs, nil
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   // in case the file is missing, we want the proper path
   creds, err := credentials(filepath.Join(home, ".git-credentials"))
   if err != nil {
      panic(err)
   }
   user := creds[0].User
   if len(os.Args) == 2 {
      prompt := os.Args[1]
      switch {
      case strings.HasPrefix(prompt, "Username"):
         fmt.Fprintln(os.Stderr, "Username")
         fmt.Println(user.Username())
      case strings.HasPrefix(prompt, "Password"):
         fmt.Fprintln(os.Stderr, "Password")
         pass, ok := user.Password()
         if ok {
            fmt.Println(pass)
         }
      }
   }
}
