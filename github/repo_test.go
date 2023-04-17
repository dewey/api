package github

import (
   "fmt"
   "testing"
   "time"
)

var repos = []repository{
   {
      name: "umber",
      homepage: "https://2e.pages.dev/umber",
   },
   {
      description: "Download APK from Google Play or send API requests",
      homepage: "https://godocs.io/2a.pages.dev/googleplay",
      name: "googleplay",
      topics: []string{
         "android",
         "google-play",
      },
   },
   {
      name: "mech",
      description: "Download media or send API requests",
      homepage: "https://godocs.io/2a.pages.dev/mech",
      topics: []string{
         "amc",
         "bandcamp",
         "cbc-gem",
         "nbc",
         "paramount",
         "roku",
         "soundcloud",
         "widevine",
         "youtube",
      },
   },
   {
      name: "nursery",
      homepage: "https://godocs.io/2a.pages.dev/nursery",
   },
   {
      name: "rosso",
      description: "Data parsers and formatters",
      homepage: "https://godocs.io/2a.pages.dev/rosso",
      topics: []string{
         "dash",
         "hls",
         "ja3",
         "json",
         "mp4",
         "protobuf",
         "xml",
      },
   },
   {
      name: "sophia",
      description: "Download media or send API requests",
   },
}

func Test_Description(t *testing.T) {
   for _, repo := range repos {
      res, err := repo.set_description()
      if err != nil {
         t.Fatal(err)
      }
      if err := res.Body.Close(); err != nil {
         t.Fatal(err)
      }
      fmt.Println(repo.name, res.Status)
      time.Sleep(time.Second)
   }
}

func Test_Topics(t *testing.T) {
   for _, repo := range repos {
      res, err := repo.set_topics()
      if err != nil {
         t.Fatal(err)
      }
      if err := res.Body.Close(); err != nil {
         t.Fatal(err)
      }
      fmt.Println(repo.name, res.Status)
      time.Sleep(time.Second)
   }
}
