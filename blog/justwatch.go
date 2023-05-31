package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "apis.justwatch.com"
   req.URL.Path = "/graphql"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(req_body)
   req.Header["Content-Type"] = []string{"application/json"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}

var req_body = strings.NewReader(`{
  "operationName": "GetTitleOffers",
  "variables": {
    "platform": "WEB",
    "nodeId": "tse371404",
    "country": "US",
    "language": "en",
    "filterBuy": {
      "monetizationTypes": [
        "BUY"
      ],
      "bestOnly": true
    },
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
    "filterRent": {
      "monetizationTypes": [
        "RENT"
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
  },
  "query": "query GetTitleOffers($nodeId: ID!, $country: Country!, $language: Language!, $filterFlatrate: OfferFilter!, $filterBuy: OfferFilter!, $filterRent: OfferFilter!, $filterFree: OfferFilter!, $platform: Platform! = WEB) {\n  node(id: $nodeId) {\n    id\n    __typename\n    ... on MovieOrShowOrSeasonOrEpisode {\n      offerCount(country: $country, platform: $platform)\n      flatrate: offers(\n        country: $country\n        platform: $platform\n        filter: $filterFlatrate\n      ) {\n        ...TitleOffer\n        __typename\n      }\n      buy: offers(country: $country, platform: $platform, filter: $filterBuy) {\n        ...TitleOffer\n        __typename\n      }\n      rent: offers(country: $country, platform: $platform, filter: $filterRent) {\n        ...TitleOffer\n        __typename\n      }\n      free: offers(country: $country, platform: $platform, filter: $filterFree) {\n        ...TitleOffer\n        __typename\n      }\n      __typename\n    }\n  }\n}\n\nfragment TitleOffer on Offer {\n  id\n  presentationType\n  monetizationType\n  retailPrice(language: $language)\n  retailPriceValue\n  currency\n  lastChangeRetailPriceValue\n  type\n  package {\n    id\n    packageId\n    clearName\n    __typename\n  }\n  standardWebURL\n  elementCount\n  availableTo\n  deeplinkRoku: deeplinkURL(platform: ROKU_OS)\n  __typename\n}\n"
}
`)
