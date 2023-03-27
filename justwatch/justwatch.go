package justwatch

import (
   "2a.pages.dev/rosso/http"
   "bytes"
   "encoding/json"
   "strings"
)

// cant use encoding.TextMarshaler because we are JSON marshalling this
func (v Variables) Text() (string, error) {
   var b strings.Builder
   country, err := get_country(v.Country_Code)
   if err != nil {
      return "", err
   }
   b.WriteString(country)
   b.WriteByte(' ')
   b.WriteString(v.Full_Path)
   return b.String(), nil
}

type details_request struct {
   Query string
   Variables Variables
}

const query = `
query GetUrlTitleDetails(
   $fullPath: String!
   $country: Country!
   $platform: Platform! = WEB
) {
   url(fullPath: $fullPath) {
      node {
         ... on MovieOrShowOrSeason {
            offers(country: $country, platform: $platform) {
               monetizationType
               presentationType
               standardWebURL
            }
         }
      }
   }
}
`

func graphQL_compact(s string) string {
   old_new := []string{
      "\n", "",
      strings.Repeat(" ", 12), " ",
      strings.Repeat(" ", 9), " ",
      strings.Repeat(" ", 6), " ",
      strings.Repeat(" ", 3), " ",
   }
   return strings.NewReplacer(old_new...).Replace(s)
}

type URLs struct {
   Href_Lang_Tags []Lang_Tag
}

// I am including `presentationType` to differentiate the different options,
// but the data seems to be incorrect in some cases. For example, JustWatch
// reports this as SD:
// fetchtv.com.au/movie/details/19285
// when the site itself reports as HD.
type Details struct {
   Data struct {
      URL struct {
         Node struct {
            Offers []struct {
               Monetization_Type string `json:"monetizationType"`
               Presentation_Type string `json:"presentationType"`
               Standard_Web_URL string `json:"standardWebURL"`
            }
         }
      }
   }
}

type Lang_Tag struct {
   Href string // fullPath
   Href_Lang string // country
}

func (t Lang_Tag) Country_Code() string {
   _, code, _ := strings.Cut(t.Href_Lang, "-")
   return code
}

func (t Lang_Tag) Language() string {
   lang, _, _ := strings.Cut(t.Href_Lang, "-")
   return lang
}

func (t Lang_Tag) Variables() Variables {
   var v Variables
   v.Country_Code = t.Country_Code()
   v.Full_Path = t.Href
   return v
}

type Variables struct {
   Country_Code string `json:"country"`
   Full_Path string `json:"fullPath"`
}

func New_URLs(c http.Client, path string) (*URLs, error) {
   req, err := http.NewRequest(
      "GET", "https://apis.justwatch.com/content/urls", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "path=" + path
   res, err := c.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   content := new(URLs)
   if err := json.NewDecoder(res.Body).Decode(content); err != nil {
      return nil, err
   }
   return content, nil
}

func (v Variables) Details(c http.Client) (*Details, error) {
   var r details_request
   r.Query = graphQL_compact(query)
   r.Variables = v
   req_body, err := json.Marshal(r)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://apis.justwatch.com/graphql", bytes.NewReader(req_body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   res, err := c.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   detail := new(Details)
   if err := json.NewDecoder(res.Body).Decode(detail); err != nil {
      return nil, err
   }
   return detail, nil
}
