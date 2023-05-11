package github

import (
   "bytes"
   "encoding/json"
   "net/http"
   "os"
)

type user struct {
   bio string
   company string
   location string
   name string
   website string
}

// REQUIRES USER SCOPE
// docs.github.com/rest/users/users#update-the-authenticated-user
func (u user) update() (*http.Response, error) {
   // Body
   body, err := json.MarshalIndent(map[string]string{
      "bio": u.bio,
      "blog": u.website,
      "company": u.company,
      "location": u.location,
      "name": u.name,
   }, "", " ")
   if err != nil {
      return nil, err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   cred, err := credential(home + "/Documents/github.json")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "PATCH",
      "https://api.github.com/user",
      bytes.NewReader(body),
   )
   req.SetBasicAuth(cred["username"], cred["password"])
   return new(http.Transport).RoundTrip(req)
}
