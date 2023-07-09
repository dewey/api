package main

import (
   "fmt"
   "golang.org/x/exp/slices"
   "net/http"
   "net/url"
   "time"
)

const digits = "abcdefghikmnopqrstuvwxz"

func main() {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Method = "HEAD"
   req.URL = new(url.URL)
   req.URL.Host = "github.com"
   req.URL.Scheme = "https"
   for _, a := range digits {
      for _, b := range digits {
         for _, c := range digits {
            for _, d := range digits {
               after := []rune{a, b, c, d}
               before := string(after)
               if before <= "abfn" {
                  continue
               }
               slices.Sort(after)
               after = slices.Compact(after)
               if len(after) <= 3 {
                  continue
               }
               req.URL.Path = "/" + before
               res, err := new(http.Transport).RoundTrip(req)
               if err != nil {
                  panic(err)
               }
               if err := res.Body.Close(); err != nil {
                  panic(err)
               }
               fmt.Println(res.Status, before)
               switch res.StatusCode {
               case http.StatusMovedPermanently:
                  // do nothing
               case http.StatusNotFound:
                  return
               case http.StatusOK:
                  // do nothing
               default:
                  panic(res.Status)
               }
               time.Sleep(399 * time.Millisecond)
            }
         }
      }
   }
}
