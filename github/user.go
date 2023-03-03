package github

import (
   "2a.pages.dev/rosso/http"
   "bytes"
   "encoding/json"
   "os"
   "strings"
)

type user struct {
   bio string
   company string
   location string
   name string
   website string
}

// docs.github.com/rest/users/users#update-the-authenticated-user
func (u user) update() (*http.Response, error) {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   creds, err := credentials(home + "/.git-credentials")
   if err != nil {
      return nil, err
   }
   user := creds[0].User
   var ref strings.Builder
   ref.WriteString("https://api.github.com/user")
   body, err := json.Marshal(map[string]string{
      "bio": u.bio,
      "blog": u.website,
      "company": u.company,
      "location": u.location,
      "name": u.name,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("PATCH", ref.String(), bytes.NewReader(body))
   if err != nil {
      return nil, err
   }
   password, ok := user.Password()
   if ok {
      req.SetBasicAuth(user.Username(), password)
   }
   /*
   -H "Accept: application/vnd.github+json" \
   -H "X-GitHub-Api-Version: 2022-11-28" \
   */
   return client.Do(req)
}
