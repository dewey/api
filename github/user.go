package github

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
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
   // URL
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   creds, err := credentials(home + "/.git-credentials")
   if err != nil {
      return nil, err
   }
   cred := creds[0].User
   req := http.Patch()
   req.Body_Bytes(body)
   if password, ok := cred.Password(); ok {
      req.SetBasicAuth(cred.Username(), password)
   }
   req.URL.Host = "api.github.com"
   req.URL.Path = "/user"
   req.URL.Scheme = "https"
   return http.Default_Client.Do(req)
}
