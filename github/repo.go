package github

import (
   "2a.pages.dev/rosso/http"
   "bufio"
   "encoding/json"
   "net/url"
   "os"
)

func (r repository) set_description() (*http.Response, error) {
   // Body
   body, err := json.Marshal(map[string]string{
      "description": r.description,
      "homepage": r.homepage,
   })
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
   user := creds[0].User
   req := http.Patch()
   req.Body_Bytes(body)
   if password, ok := user.Password(); ok {
      req.SetBasicAuth(user.Username(), password)
   }
   req.URL.Host = "api.github.com"
   req.URL.Path = "/repos/" + user.Username() + "/" + r.name
   req.URL.Scheme = "https"
   return http.Default_Client.Do(req)
}

func credentials(name string) ([]url.URL, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   var refs []url.URL
   buf := bufio.NewScanner(file)
   for buf.Scan() {
      var ref url.URL
      err := ref.UnmarshalBinary(buf.Bytes())
      if err != nil {
         return nil, err
      }
      refs = append(refs, ref)
   }
   return refs, nil
}

func (r repository) set_topics() (*http.Response, error) {
   // Body
   body, err := json.Marshal(map[string][]string{
      "names": r.topics,
   })
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
   user := creds[0].User
   req := http.Put()
   req.Body_Bytes(body)
   if password, ok := user.Password(); ok {
      req.SetBasicAuth(user.Username(), password)
   }
   req.URL.Host = "api.github.com"
   req.URL.Path = "/repos/" + user.Username() + "/" + r.name + "/topics"
   req.URL.Scheme = "https"
   return http.Default_Client.Do(req)
}

type repository struct {
   topics []string
   name string
   description string
   homepage string
}

