package github

import (
   "bytes"
   "encoding/json"
   "net/http"
   "os"
)

func credential(name string) (map[string]string, error) {
   data, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   var cred map[string]string
   if err := json.Unmarshal(data, &cred); err != nil {
      return nil, err
   }
   return cred, nil
}

func (r repository) set_description() (*http.Response, error) {
   // Body
   body, err := json.Marshal(map[string]string{
      "description": r.description,
      "homepage": r.homepage,
   })
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
      "PATCH", "https://api.github.com/repos/" + cred["username"] + "/" + r.name,
      bytes.NewReader(body),
   )
   req.SetBasicAuth(cred["username"], cred["password"])
   return new(http.Transport).RoundTrip(req)
}

func (r repository) set_topics() (*http.Response, error) {
   body, err := json.Marshal(map[string][]string{
      "names": r.topics,
   })
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
      "PUT",
      "https://api.github.com/repos/" + cred["username"] + "/" + r.name + "/topics",
      bytes.NewReader(body),
   )
   req.SetBasicAuth(cred["username"], cred["password"])
   return new(http.Transport).RoundTrip(req)
}

type repository struct {
   topics []string
   name string
   description string
   homepage string
}
