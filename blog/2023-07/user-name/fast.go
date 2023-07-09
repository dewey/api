package main

import (
   "bufio"
   "fmt"
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
   file, err := os.Open("four.txt")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   buf := bufio.NewScanner(file)
   for buf.Scan() {
      text := buf.Text()
      if text <= "blog" {
         continue
      }
      req.URL.Path = "/" + text
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         panic(err)
      }
      if err := res.Body.Close(); err != nil {
         panic(err)
      }
      fmt.Println(res.Status, text)
      switch res.StatusCode {
      case http.StatusMovedPermanently:
         // do nothing
      case http.StatusNotFound:
         file, err := os.Create(text)
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
