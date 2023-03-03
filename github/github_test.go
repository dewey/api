package github

import (
   "fmt"
   "testing"
   "time"
)

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

var repos = []repository{
   {
      name: "git",
      description: "Git implementation",
   },
   {
      name: "googleplay",
      description: "Download APK from Google Play or send API requests",
      topics: []string{
         "android",
         "google-play",
      },
   },
   {
      name: "mech",
      description: "Download media or send API requests",
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
      name: "rosso",
      description: "Data parsers and formatters",
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
