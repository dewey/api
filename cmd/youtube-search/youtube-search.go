package main

import (
   "flag"
   "net/url"
   "os/exec"
   "strconv"
)

func (f flags) address() string {
   var b []byte
   b = append(b, "intext:"...)
   b = strconv.AppendQuote(b, f.artist + " topic")
   b = append(b, " intitle:"...)
   b = strconv.AppendQuote(b, f.song)
   var ref url.URL
   ref.Scheme = "https"
   ref.Host = "www.youtube.com"
   ref.Path = "results"
   ref.RawQuery = "search_query=" + url.QueryEscape(string(b))
   return ref.String()
}

type flags struct {
   artist string
   browser string
   song string
}

func main() {
   var f flags
   // a
   flag.StringVar(&f.artist, "a", "", "artist")
   // b
   flag.StringVar(
      &f.browser, "b", `C:\Program Files\Mozilla Firefox\firefox`, "browser",
   )
   // s
   flag.StringVar(&f.song, "s", "", "song")
   flag.Parse()
   if f.artist != "" && f.song != "" {
      err := exec.Command(f.browser, f.address()).Start()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
