package github

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
   "time"
)

var repos = []repository{
   {
      name: "api",
      homepage: "https://godocs.io/154.pages.dev/api",
      topics: []string{
         "github",
         "justwatch",
         "mullvad",
         "musicbrainz",
      },
   },
   {
      name: "encoding",
      description: "Data parsers and formatters",
      topics: []string{
         "dash",
         "hls",
         "json",
         "mp4",
         "protobuf",
         "xml",
      },
      homepage: "https://godocs.io/154.pages.dev/encoding",
   },
   {
      description: "Download APK from Google Play or send API requests",
      name: "google",
      topics: []string{"android"},
      homepage: "https://godocs.io/154.pages.dev/google",
   },
   {
      name: "media",
      description: "Download media or send API requests",
      topics: []string{
         "amc",
         "bandcamp",
         "cbc-gem",
         "nbc",
         "paramount",
         "roku",
         "soundcloud",
         "twitter",
         "youtube",
      },
      homepage: "https://godocs.io/154.pages.dev/media",
   },
   {
      name: "strconv",
      homepage: "https://godocs.io/154.pages.dev/strconv",
   },
   {
      description: "low-level access to the ClientHello for mimicry purposes",
      name: "tls",
      topics: []string{
         "android",
         "censorship-circumvention",
         "crypto",
      },    
      homepage: "https://godocs.io/154.pages.dev/tls",
   },
   {
      name: "umber",
      homepage: "https://154.pages.dev/umber",
   },
   {
      name: "widevine",
      description: "DRM",
      homepage: "https://godocs.io/154.pages.dev/widevine",
   },
}
func Test_Description(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := user_info(home + "/github.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, repo := range repos {
      fmt.Println(repo.name)
      err := repo.set_description(u)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

func user_info(name string) (map[string]string, error) {
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

func Test_Topics(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := user_info(home + "/github.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, repo := range repos {
      if repo.topics != nil {
         fmt.Println(repo.name)
         err := repo.set_topics(u)
         if err != nil {
            t.Fatal(err)
         }
         time.Sleep(time.Second)
      }
   }
}

func Test_User(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   m, err := user_info(home + "/github.json")
   if err != nil {
      t.Fatal(err)
   }
   u := user{
      bio: "email srpen6@gmail.com, Discord srpen6",
      company: "looking for work",
      location: "Dallas",
      name: "Steven Penny",
      website: "https://discord.com/invite/WWq6rFb8Rf",
   }
   if err := u.update(m); err != nil {
      t.Fatal(err)
   }
}

