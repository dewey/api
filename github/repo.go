package github

import (
   "2a.pages.dev/rosso/http"
   "bufio"
   "bytes"
   "encoding/json"
   "net/url"
   "os"
   "strings"
)

func (r repository) set_description() (*http.Response, error) {
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
   ref.WriteString("https://api.github.com/repos/")
   ref.WriteString(user.Username())
   ref.WriteByte('/')
   ref.WriteString(r.name)
   body, err := json.Marshal(map[string]string{
      "description": r.description,
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
   return client.Do(req)
}

var client = http.Default_Client

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

type repository struct {
   name string
   description string
   topics []string
}

func (r repository) set_topics() (*http.Response, error) {
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
   ref.WriteString("https://api.github.com/repos/")
   ref.WriteString(user.Username())
   ref.WriteByte('/')
   ref.WriteString(r.name)
   ref.WriteString("/topics")
   body, err := json.Marshal(map[string][]string{
      "names": r.topics,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("PUT", ref.String(), bytes.NewReader(body))
   if err != nil {
      return nil, err
   }
   password, ok := user.Password()
   if ok {
      req.SetBasicAuth(user.Username(), password)
   }
   return client.Do(req)
}
