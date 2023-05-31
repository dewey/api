package main

import (
   "bytes"
   "io"
   "net/http"
   "net/url"
   "strings"
   "fmt"
)

const query = `
query GetTitleOffers(
   $country: Country!
   $nodeId: ID!
   $platform: Platform! = WEB
) {
   node(id: $nodeId) {
      ... on MovieOrShowOrSeasonOrEpisode {
         offers(
            country: $country
            platform: $platform
         ) {
            ... on Offer {
               standardWebURL
            }
         }
      }
   }
}
`

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "apis.justwatch.com"
   req.URL.Path = "/graphql"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/json"}
   req_body := fmt.Sprintf(`
{
   "query": %q,
   "variables": {
      "nodeId": "tse371404",
      "country": "US",
      "filterFlatrate": {
         "monetizationTypes": [
            "FLATRATE",
            "FLATRATE_AND_BUY",
            "ADS",
            "FREE",
            "CINEMA"
         ],
         "bestOnly": true
      },
      "filterFree": {
         "monetizationTypes": [
            "ADS",
            "FREE"
         ],
         "bestOnly": true
      }
   }
}
`, query)
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
   // amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152
   if bytes.Contains(res_body, []byte("/season-1-instinct--1011152")) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}
