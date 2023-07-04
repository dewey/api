package github

import (
   "fmt"
   "testing"
   "time"
)

func Test_Topics(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := user(home + "/github.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, repo := range repos {
      if repo.topics != nil {
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
}

func Test_User(t *testing.T) {
   u := user{
      bio: "email srpen6@gmail.com, Discord srpen6",
      company: "looking for work",
      location: "Dallas",
      name: "Steven Penny",
      website: "https://discord.com/invite/WWq6rFb8Rf",
   }
   res, err := u.update()
   if err != nil {
      t.Fatal(err)
   }
   if err := res.Body.Close(); err != nil {
      t.Fatal(err)
   }
   fmt.Println(res.Status)
}

func user(name string) (map[string]string, error) {
   b, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   var m map[string]string
   if err := json.Unmarshal(b, &m); err != nil {
      return nil, err
   }
   return m, nil
}

var repos = []repository{
   {
      name: "mech",
      description: "Download media or send API requests",
      homepage: "https://godocs.io/2a.pages.dev/mech",
      topics: []string{
         // justwatch.com/us/provider/amc-plus
         "amc-plus",
         // bandcamp.com
         "bandcamp",
         // justwatch.com/ca/provider/cbc-gem
         "cbc-gem",
         // justwatch.com/us/provider/nbc
         "nbc",
         // justwatch.com/us/provider/paramount-plus
         "paramount-plus",
         // soundcloud.com
         "soundcloud",
         // justwatch.com/us/provider/the-roku-channel
         "roku-channel",
         // twitter.com
         "twitter",
         // widevine.com
         "widevine",
         // youtube.com
         "youtube",
      },
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
      description: "Cisco Android and web authentication",
      homepage: "https://godocs.io/2a.pages.dev/meraki",
      name: "meraki",
      topics: []string{"2fa"},
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
         "json",
         "mp4",
         "protobuf",
         "xml",
      },
   },
   {
      description: "low-level access to the ClientHello for mimicry purposes",
      homepage: "https://godocs.io/2a.pages.dev/tls",
      name: "tls",
      topics: []string{
         "android",
         "censorship-circumvention",
         "crypto",
      } ,    
   },
   {
      name: "umber",
      homepage: "https://2e.pages.dev/umber",
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
