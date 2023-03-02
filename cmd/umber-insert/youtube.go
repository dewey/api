package main

import (
   "2a.pages.dev/mech/youtube"
   "flag"
   "fmt"
   "net/http"
   "net/url"
   "path"
   "sort"
   "strconv"
   "strings"
   "time"
)

func get_image(video_ID string) (string, bool) {
   var imgs []youtube.Image
   for _, img := range youtube.Images {
      if img.Height < 720 {
         imgs = append(imgs, img)
      }
   }
   sort.SliceStable(imgs, func(a, b int) bool {
      com := imgs[a].Height - imgs[b].Height
      if com != 0 {
         return com >= 1
      }
      def := func(i int) int {
         return strings.Index(imgs[i].Name, "default")
      }
      com = def(a) - def(b)
      if com != 0 {
         return com >= 1
      }
      def = func(i int) int {
         return strings.Index(imgs[i].Name, "webp")
      }
      return def(b) < def(a)
   })
   for key, val := range imgs {
      ref := val.Address(video_ID)
      fmt.Println("HEAD", ref)
      res, err := http.Head(ref)
      if err == nil && res.StatusCode == http.StatusOK {
         if key == 0 {
            return "", false
         }
         return ref, true
      }
   }
   return "", false
}
type youtube_set struct {
   *flag.FlagSet
   video_ID string
}

func new_youtube() *youtube_set {
   var set youtube_set
   set.FlagSet = flag.NewFlagSet("youtube", flag.ExitOnError)
   set.StringVar(&set.video_ID, "b", "", "video ID")
   set.Func("a", "address", func(s string) error {
      return youtube.Video_ID(s, &set.video_ID)
   })
   return &set
}

func (y *youtube_set) parse(arg []string) (*record, error) {
   y.Parse(arg)
   val := make(url.Values)
   now := strconv.FormatInt(time.Now().Unix(), 36)
   val.Set("a", now)
   val.Set("p", "y")
   val.Set("b", y.video_ID)
   ref, ok := get_image(y.video_ID)
   if ok {
      val.Set("c", path.Base(ref))
   }
   play, err := youtube.Mobile_Web().Player(y.video_ID, nil)
   if err != nil {
      return nil, err
   }
   var rec record
   rec.S = play.Video_Details.Author + " - " + play.Video_Details.Title
   fmt.Println(play.Video_Details.Short_Description)
   year, _, ok := strings.Cut(play.Publish_Date(), "-")
   if ok {
      val.Set("y", year)
   }
   rec.Q = val.Encode()
   return &rec, nil
}
