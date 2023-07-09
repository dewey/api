package main

import (
   "fmt"
   "golang.org/x/exp/slices"
   "net/http"
   "net/url"
   "os"
   "time"
)

func main() {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Method = "HEAD"
   req.URL = new(url.URL)
   req.URL.Host = "github.com"
   req.URL.Scheme = "https"
   i := 999
   for {
      i++
      if i <= 1092 {
         continue
      }
      before := fmt.Sprint(i)
      after := []byte(before)
      slices.Sort(after)
      after = slices.Compact(after)
      if len(after) < len(before) {
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
         file, err := os.Create(before)
         if err != nil {
            panic(err)
         }
         if err := file.Close(); err != nil {
            panic(err)
         }
      case http.StatusOK:
         // do nothing
      default:
         panic(res.Status)
      }
      time.Sleep(399 * time.Millisecond)
   }
}