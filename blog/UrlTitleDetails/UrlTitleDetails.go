package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "apis.justwatch.com"
   req.URL.Path = "/graphql"
   req.URL.Scheme = "https"
   req.Method = "POST"
   req.Header["Content-Type"] = []string{"application/json"}
   //full_path := "/us/tv-show/orphan-black/season-1"
   full_path := "/us/tv-show/orphan-black"
   req_body := fmt.Sprintf(`
   {
     "operationName": "GetUrlTitleDetails",
     "query": %q,
     "variables": {
       "platform": "WEB",
       "language": "en",
       "country": "US",
       "episodeMaxLimit": 20,
       "fullPath": %q
     }
   }
   `, query, full_path)
   req.Body = io.NopCloser(strings.NewReader(req_body))
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   fmt.Println(string(res_body))
   if bytes.Contains(res_body, []byte("tse371404")) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}

const query = `
query GetUrlTitleDetails($fullPath: String!) {
  url(fullPath: $fullPath) {
    node {
      ... on Season {
        episodes() {
          id
        }
      }
      ... on Show {
        episodes() {
          id
        }
      }
    }
  }
}
`
