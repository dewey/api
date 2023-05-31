package justwatch

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

func compact(offers []offer) map[string]struct{} {
   m := make(map[string]struct{})
   for _, offer := range offers {
      m[offer.Standard_Web_URL] = struct{}{}
   }
   return m
}

type offer struct {
   Standard_Web_URL string `json:"standardWebUrl"`
}

func title_offers(node_ID string) ([]offer, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Query string `json:"query"`
         Variables struct {
            Country string `json:"country"`
            Node_ID string `json:"nodeId"`
         } `json:"variables"`
      }
      s.Query = query
      s.Variables.Country = "US"
      s.Variables.Node_ID = node_ID
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://apis.justwatch.com/graphql", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var s struct {
      Data struct {
         Node struct {
            Offers []offer
         }
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   if err := res.Body.Close(); err != nil {
      return nil, err
   }
   return s.Data.Node.Offers, nil
}

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

