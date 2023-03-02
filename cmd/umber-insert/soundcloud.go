package main

import (
   "2a.pages.dev/mech/soundcloud"
   "flag"
   "net/url"
   "path"
   "strconv"
   "strings"
   "time"
)

type soundcloud_set struct {
   *flag.FlagSet
   address string
}

func new_soundcloud() *soundcloud_set {
   var set soundcloud_set
   set.FlagSet = flag.NewFlagSet("soundcloud", flag.ExitOnError)
   set.StringVar(&set.address, "a", "", "address")
   return &set
}

func (s *soundcloud_set) parse(arg []string) (*record, error) {
   s.Parse(arg)
   val := make(url.Values)
   now := strconv.FormatInt(time.Now().Unix(), 36)
   val.Set("a", now)
   val.Set("p", "s")
   tracks, err := soundcloud.Resolve(s.address)
   if err != nil {
      return nil, err
   }
   track := tracks[0]
   var row record
   row.S = track.Title
   val.Set("b", strconv.FormatInt(track.ID, 10))
   val.Set("c", path.Base(track.Artwork()))
   year, _, ok := strings.Cut(track.Display_Date, "-")
   if ok {
      val.Set("y", year)
   }
   row.Q = val.Encode()
   return &row, nil
}
