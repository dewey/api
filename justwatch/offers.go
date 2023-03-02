package justwatch

import (
   "fmt"
   "sort"
   "strings"
)

// lets match Variables type
func (o Offers) Text() (string, error) {
   var b strings.Builder
   for mon_type, url_codes := range o {
      if b.Len() >= 1 {
         b.WriteByte('\n')
      }
      b.WriteString(mon_type)
      for _, web_URL := range sort_keys(url_codes) {
         b.WriteByte('\n')
         b.WriteString(tab)
         b.WriteString(web_URL)
         for _, code := range sort_keys(url_codes[web_URL]) {
            b.WriteByte('\n')
            b.WriteString(tab)
            b.WriteString("- ")
            country, err := get_country(code)
            if err != nil {
               return "", err
            }
            b.WriteString(country)
         }
      }
   }
   return b.String(), nil
}

type Country_Codes map[string]struct{}

// map[monetizationType]map[standardWebURL]Country_Codes
type Offers map[string]map[string]Country_Codes

func sort_keys[M ~map[string]V, V any](group M) []string {
   var keys []string
   for key := range group {
      keys = append(keys, key)
   }
   sort.Strings(keys)
   return keys
}

var countries = map[string]string{
   "AG": "Antigua and Barbuda",
   "AU": "Australia",
   "BB": "Barbados",
   "BM": "Bermuda",
   "BS": "Bahamas",
   "CA": "Canada",
   "GB": "United Kingdom",
   "DK": "Denmark",
   "FJ": "Fiji",
   "GG": "Guernsey",
   "GH": "Ghana",
   "GI": "Gibraltar",
   "ID": "Indonesia",
   "IE": "Ireland",
   "IN": "India",
   "JM": "Jamaica",
   "KE": "Kenya",
   "LC": "Saint Lucia",
   "MY": "Malaysia",
   "NG": "Nigeria",
   "NL": "Netherlands",
   "NO": "Norway",
   "NZ": "New Zealand",
   "PH": "Philippines",
   "SG": "Singapore",
   "TC": "Turks and Caicos",
   "TH": "Thailand",
   "TT": "Trinidad and Tobago",
   "UG": "Uganda",
   "US": "United States",
   "ZA": "South Africa",
   "ZM": "Zambia",
}

func get_country(code string) (string, error) {
   country := countries[code]
   if country == "" {
      return "", fmt.Errorf("invalid country code %q", code)
   }
   return country, nil
}

func (o Offers) Add(country_code string, detail *Details) {
   for _, node := range detail.Data.URL.Node.Offers {
      offer := o[node.Monetization_Type]
      if offer == nil {
         offer = make(map[string]Country_Codes)
      }
      codes := offer[node.Standard_Web_URL]
      if codes == nil {
         codes = make(Country_Codes)
      }
      codes[country_code] = struct{}{}
      offer[node.Standard_Web_URL] = codes
      o[node.Monetization_Type] = offer
   }
}

func (o Offers) Stream() Offers {
   p := make(Offers)
   for m_type, offer := range o {
      if !buy_rent[m_type] {
         p[m_type] = offer
      }
   }
   return p
}

const tab = "   "

var buy_rent = map[string]bool{
   "BUY": true,
   "RENT": true,
}
