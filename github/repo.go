package github

import (
   "bytes"
   "encoding/json"
   "net/http"
   "os"
   "strings"
)

func user_info(name string) ([]string, error) {
   text, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   return strings.Split(string(text), "\n"), nil
}

func (r repository) set_description() (*http.Response, error) {
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
   user, err := user_info(home + "/Documents/github.txt")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "PATCH", "https://api.github.com/repos/" + user[0] + "/" + r.name,
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.SetBasicAuth(user[0], user[1])
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
   user, err := user_info(home + "/Documents/github.txt")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "PUT",
      "https://api.github.com/repos/" + user[0] + "/" + r.name + "/topics",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.SetBasicAuth(user[0], user[1])
   return new(http.Transport).RoundTrip(req)
}

type repository struct {
   topics []string
   name string
   description string
   homepage string
}
