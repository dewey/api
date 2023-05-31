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
          package {
            id
            packageId
            __typename
          }
          __typename
        }
        promotedBundles(country: $country, platform: $platform) {
          promotionUrl
          __typename
        }
        availableTo(country: $country, platform: $platform) {
          availableCountDown(country: $country)
          availableToDate
          package {
            id
            shortName
            __typename
          }
          __typename
        }
        fallBackClips: content(country: "US", language: "en") {
          videobusterClips: clips(providers: [VIDEOBUSTER]) {
            externalId
            provider
            __typename
          }
          dailymotionClips: clips(providers: [DAILYMOTION]) {
            externalId
            provider
            __typename
          }
          __typename
        }
        content(country: $country, language: $language) {
          backdrops {
            backdropUrl
            __typename
          }
          clips {
            externalId
            provider
            __typename
          }
          videobusterClips: clips(providers: [VIDEOBUSTER]) {
            externalId
            provider
            __typename
          }
          dailymotionClips: clips(providers: [DAILYMOTION]) {
            externalId
            provider
            __typename
          }
          videobusterClips: clips(providers: [VIDEOBUSTER]) {
            externalId
            __typename
          }
          externalIds {
            imdbId
            __typename
          }
          fullPath
          genres {
            shortName
            __typename
          }
          posterUrl
          runtime
          isReleased
          scoring {
            imdbScore
            imdbVotes
            tmdbPopularity
            tmdbScore
            __typename
          }
          shortDescription
          title
          originalReleaseYear
          originalReleaseDate
          upcomingReleases(releaseTypes: DIGITAL) {
            releaseCountDown(country: $country)
            releaseDate
            label
            package {
              id
              packageId
              shortName
              __typename
            }
            __typename
          }
          ... on MovieOrShowContent {
            originalTitle
            ageCertification
            credits {
              role
              name
              characterName
              personId
              __typename
            }
            productionCountries
            __typename
          }
          ... on SeasonContent {
            seasonNumber
            __typename
          }
          __typename
        }
        __typename
      }
      ... on MovieOrShow {
        watchlistEntry {
          createdAt
          __typename
        }
        likelistEntry {
          createdAt
          __typename
        }
        dislikelistEntry {
          createdAt
          __typename
        }
        customlistEntries {
          createdAt
          genericTitleList {
            id
            __typename
          }
          __typename
        }
        __typename
      }
      ... on Movie {
        permanentAudiences
        seenlistEntry {
          createdAt
          __typename
        }
        __typename
      }
      ... on Show {
        permanentAudiences
        totalSeasonCount
        seenState(country: $country) {
          progress
          seenEpisodeCount
          __typename
        }
        seasons(sortDirection: DESC) {
          id
          objectId
          objectType
          availableTo(country: $country, platform: $platform) {
            availableToDate
            availableCountDown(country: $country)
            package {
              id
              shortName
              __typename
            }
            __typename
          }
          content(country: $country, language: $language) {
            posterUrl
            seasonNumber
            fullPath
            upcomingReleases(releaseTypes: DIGITAL) {
              releaseDate
              releaseCountDown(country: $country)
              package {
                id
                shortName
                __typename
              }
              __typename
            }
            isReleased
            __typename
          }
          show {
            id
            objectId
            objectType
            watchlistEntry {
              createdAt
              __typename
            }
            content(country: $country, language: $language) {
              title
              __typename
            }
            __typename
          }
          __typename
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
            upcomingReleases {
              releaseDate
              label
              __typename
            }
            __typename
          }
          seenlistEntry {
            createdAt
            __typename
          }
          __typename
        }
        __typename
      }
      ... on Season {
        totalEpisodeCount
        episodes(limit: $episodeMaxLimit) {
          id
          objectType
          objectId
          seenlistEntry {
            createdAt
            __typename
          }
          content(country: $country, language: $language) {
            title
            shortDescription
            episodeNumber
            seasonNumber
            isReleased
            upcomingReleases(releaseTypes: DIGITAL) {
              releaseDate
              label
              package {
                id
                packageId
                __typename
              }
              __typename
            }
            __typename
          }
          __typename
        }
        show {
          id
          objectId
          objectType
          totalSeasonCount
          customlistEntries {
            createdAt
            genericTitleList {
              id
              __typename
            }
            __typename
          }
          fallBackClips: content(country: "US", language: "en") {
            videobusterClips: clips(providers: [VIDEOBUSTER]) {
              externalId
              provider
              __typename
            }
            dailymotionClips: clips(providers: [DAILYMOTION]) {
              externalId
              provider
              __typename
            }
            __typename
          }
          content(country: $country, language: $language) {
            title
            ageCertification
            fullPath
            credits {
              role
              name
              characterName
              personId
              __typename
            }
            productionCountries
            externalIds {
              imdbId
              __typename
            }
            upcomingReleases(releaseTypes: DIGITAL) {
              releaseDate
              __typename
            }
            backdrops {
              backdropUrl
              __typename
            }
            posterUrl
            isReleased
            videobusterClips: clips(providers: [VIDEOBUSTER]) {
              externalId
              provider
              __typename
            }
            dailymotionClips: clips(providers: [DAILYMOTION]) {
              externalId
              provider
              __typename
            }
            __typename
          }
          seenState(country: $country) {
            progress
            __typename
          }
          watchlistEntry {
            createdAt
            __typename
          }
          dislikelistEntry {
            createdAt
            __typename
          }
          likelistEntry {
            createdAt
            __typename
          }
          __typename
        }
        seenState(country: $country) {
          progress
          __typename
        }
        __typename
      }
      __typename
    }
    __typename
  }
}
`
