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
            }
         }
      }
   }
}

func new_title_details(full_path string) (*title_details, error) {
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
