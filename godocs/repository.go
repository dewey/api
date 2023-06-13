package godocs

import (
   "encoding/xml"
   "errors"
   "net/http"
   "strings"
)

func repository(imp string) (string, error) {
   if strings.HasPrefix(imp, "github.com/") {
      return easy(imp), nil
   }
   return hard(imp)
}

func easy(imp string) string {
   split := strings.Split(imp, "/")
   return "https://" + strings.Join(split[:3], "/")
}

func hard(imp string) (string, error) {
   rep := func() string {
      split := strings.Split(imp, "/")
      return "https://" + strings.Join(split[:2], "/")
   }()
   res, err := http.Get(rep)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   var s struct {
      Head    struct {
         Meta []struct {
            Name    string `xml:"name,attr"`
            Content string `xml:"content,attr"`
         } `xml:"meta"`
      } `xml:"head"`
   }
   dec := xml.NewDecoder(res.Body)
   dec.AutoClose = xml.HTMLAutoClose
   dec.Strict = false
   if err := dec.Decode(&s); err != nil {
      return "", err
   }
   for _, meta := range s.Head.Meta {
      if meta.Name == "go-import" {
         i := strings.LastIndexByte(meta.Content, ' ')
         return meta.Content[i+1:], nil
      }
   }
   return "", errors.New("go-import not found")
}
