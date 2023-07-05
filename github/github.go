package github

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

func (r repository) set_topics(m map[string]string) error {
   body, err := json.Marshal(map[string][]string{
      "names": r.topics,
   })
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "PUT",
      "https://api.github.com/repos/" + m["username"] + "/" + r.name + "/topics",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.SetBasicAuth(m["username"], m["password"])
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return nil
}

type user struct {
   bio string
   company string
   location string
   name string
   website string
}

// REQUIRES USER SCOPE
// docs.github.com/rest/users/users#update-the-authenticated-user
func (u user) update(m map[string]string) error {
   body, err := func() ([]byte, error) {
      m := map[string]string{
         "bio": u.bio,
         "blog": u.website,
         "company": u.company,
         "location": u.location,
         "name": u.name,
      }
      return json.MarshalIndent(m, "", " ")
   }()
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "PATCH", "https://api.github.com/user", bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.SetBasicAuth(m["username"], m["password"])
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return nil
}

type repository struct {
   description string
   homepage string
   name string
   topics []string
}

func (r repository) set_description(m map[string]string) error {
   body, err := json.Marshal(map[string]string{
      "description": r.description,
      "homepage": r.homepage,
   })
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "PATCH", "https://api.github.com/repos/" + m["username"] + "/" + r.name,
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.SetBasicAuth(m["username"], m["password"])
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return nil
}
