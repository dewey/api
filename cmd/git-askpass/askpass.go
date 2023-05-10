package main

import (
   "bufio"
   "net/url"
   "os"
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
