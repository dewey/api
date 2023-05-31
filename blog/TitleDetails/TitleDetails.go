package justwatch

import (
   "encoding/json"
   "fmt"
   "net/http"
   "strings"
)

type title_details struct {
   Data struct {
      URL struct {
         Node struct {
            Episodes []struct {
               ID string
               Content struct {
                  Episode_Number int `json:"episodeNumber"`
                  Season_Number int `json:"seasonNumber"`
                  Title string
               }
            }
         }
      }
   }
}

func new_title_details(full_path string) (*title_details, error) {
   req_body := fmt.Sprintf(`
   {
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
   req, err := http.NewRequest(
      "POST", "https://apis.justwatch.com/graphql", strings.NewReader(req_body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   detail := new(title_details)
   if err := json.NewDecoder(res.Body).Decode(detail); err != nil {
      return nil, err
   }
   return detail, nil
}

const query = `
query GetUrlTitleDetails(
   $country: Country!
   $fullPath: String!
   $language: Language!
) {
  url(fullPath: $fullPath) {
    node {
      ... on Show {
        episodes() {
          content(country: $country, language: $language) {
            episodeNumber
            seasonNumber
            title
          }
          id
        }
      }
    }
  }
}
`
