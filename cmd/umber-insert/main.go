package main

import (
   "embed"
   "encoding/json"
   "net/url"
   "os"
   "path"
   "strconv"
   "strings"
   "time"
)

type inserter interface {
   Date() string
   ID() string
   Image() string
   Title() string
}

func Insert(i inserter, platform string) (*record, error) {
   val := make(url.Values)
   val.Set("a", strconv.FormatInt(time.Now().Unix(), 36))
   val.Set("b", i.ID())
   val.Set("c", path.Base(i.Image()))
   val.Set("p", platform)
   if year, _, ok := strings.Cut(i.Date(), "-"); ok {
      val.Set("y", year)
   }
   var row record
   row.Q = val.Encode()
   row.S = i.Title()
   return &row, nil
}

//go:embed insert.json
var content embed.FS

type record struct {
   Q string
   S string
}

func main() {
   buf, err := content.ReadFile("insert.json")
   if err != nil {
      panic(err)
   }
   var config struct {
      Umber string
   }
   if err := json.Unmarshal(buf, &config); err != nil {
      panic(err)
   }
   file, err := os.Open(config.Umber)
   if err != nil {
      panic(err)
   }
   defer file.Close()
   var recs []*record
   if err := json.NewDecoder(file).Decode(&recs); err != nil {
      panic(err)
   }
   if len(os.Args) >= 3 {
      arg := os.Args[2:]
      var rec *record
      switch os.Args[1] {
      case "backblaze":
         rec, err = new_backblaze().parse(arg)
      case "bandcamp":
         rec, err = new_bandcamp().parse(arg)
      case "soundcloud":
         rec, err = new_soundcloud().parse(arg)
      case "youtube":
         rec, err = new_youtube().parse(arg)
      }
      if err != nil {
         panic(err)
      }
      recs = append([]*record{rec}, recs...)
      file, err := os.Create(config.Umber)
      if err != nil {
         panic(err)
      }
      defer file.Close()
      enc := json.NewEncoder(file)
      enc.SetEscapeHTML(false)
      enc.SetIndent("", " ")
      if err := enc.Encode(recs); err != nil {
         panic(err)
      }
   } else {
      new_backblaze().Usage()
      new_bandcamp().Usage()
      new_soundcloud().Usage()
      new_youtube().Usage()
   }
}
