package main

import (
   "encoding/json"
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
   req.URL = new(url.URL)
   req.URL.Host = "dash.cloudflare.com"
   req.URL.Path = "/api/v4/accounts/a0f236930132796e8f429a7dbc35f0a4/pages/get_subdomain"
   req.URL.Scheme = "https"
   {
      h, err := os.UserHomeDir()
      if err != nil {
         panic(err)
      }
      b, err := os.ReadFile(h + "/pages.json")
      if err != nil {
         panic(err)
      }
      var s struct {
         V_Ses_2 string `json:"vses2"`
      }
      json.Unmarshal(b, &s)
      req.Header["Cookie"] = []string{"vses2=" + s.V_Ses_2}
   }
   i := -1
   for {
      i++
      before := fmt.Sprint(i)
      after := []byte(before)
      slices.Sort(after)
      after = slices.Compact(after)
      if len(after) < len(before) {
         continue
      }
      req.URL.RawQuery = "project_name=" + before
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         panic(err)
      }
      var s struct {
         Result struct {
            Subdomain string
         }
      }
      if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
         panic(err)
      }
      if err := res.Body.Close(); err != nil {
         panic(err)
      }
      fmt.Println(s.Result.Subdomain)
      if len(s.Result.Subdomain) == len(before)+len(".pages.dev") {
         break
      }
      time.Sleep(199 * time.Millisecond)
   }
}
