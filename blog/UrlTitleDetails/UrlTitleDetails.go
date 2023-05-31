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
   req_body := fmt.Sprintf(`
   {
     "operationName": "GetUrlTitleDetails",
     "variables": {
       "platform": "WEB",
       "fullPath": "/us/tv-show/orphan-black/season-1",
       "language": "en",
       "country": "US",
       "episodeMaxLimit": 20
     },
     "query": %q
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
   if bytes.Contains(res_body, []byte("tse371404")) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}

const query = `
query GetUrlTitleDetails($fullPath: String!, $country: Country!, $language: Language!, $episodeMaxLimit: Int, $platform: Platform! = WEB) {
  url(fullPath: $fullPath) {
    id
    metaDescription
    metaKeywords
    metaRobots
    metaTitle
    heading1
    heading2
    htmlContent
    node {
      id
      ... on MovieOrShowOrSeason {
        objectType
        objectId
        offerCount(country: $country, platform: $platform)
        offers(country: $country, platform: $platform) {
          monetizationType
        }
        promotedBundles(country: $country, platform: $platform) {
          promotionUrl
        }
        availableTo(country: $country, platform: $platform) {
          availableCountDown(country: $country)
          availableToDate
        }
        content(country: $country, language: $language) {
          fullPath
          posterUrl
          runtime
          isReleased
          shortDescription
          title
          originalReleaseYear
          originalReleaseDate
        }
      }
      ... on MovieOrShow {
        watchlistEntry {
          createdAt
        }
        likelistEntry {
          createdAt
        }
        dislikelistEntry {
          createdAt
        }
        customlistEntries {
          createdAt
        }
      }
      ... on Movie {
        permanentAudiences
        seenlistEntry {
          createdAt
        }
      }
      ... on Show {
        permanentAudiences
        totalSeasonCount
        seenState(country: $country) {
          progress
          seenEpisodeCount
        }
        seasons(sortDirection: DESC) {
          id
          objectId
          objectType
          show {
            id
            objectId
            objectType
          }
        }
        recentEpisodes: episodes(
          sortDirection: DESC
          limit: 3
          releasedInCountry: $country
        ) {
          id
          objectId
          content(country: $country, language: $language) {
            title
            shortDescription
            episodeNumber
            seasonNumber
            isReleased
          }
          seenlistEntry {
            createdAt
          }
        }
      }
      ... on Season {
        totalEpisodeCount
        episodes(limit: $episodeMaxLimit) {
          id
          objectType
          objectId
          seenlistEntry {
            createdAt
          }
          content(country: $country, language: $language) {
            title
            shortDescription
            episodeNumber
            seasonNumber
            isReleased
          }
        }
        show {
          id
          objectId
          objectType
          totalSeasonCount
          customlistEntries {
            createdAt
          }
          content(country: $country, language: $language) {
            title
            ageCertification
            fullPath
            productionCountries
            posterUrl
            isReleased
          }
          seenState(country: $country) {
            progress
          }
          watchlistEntry {
            createdAt
          }
          dislikelistEntry {
            createdAt
          }
          likelistEntry {
            createdAt
          }
        }
        seenState(country: $country) {
          progress
        }
      }
    }
  }
}
`
