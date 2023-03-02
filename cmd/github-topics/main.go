package main

import "flag"

func main() {
   var f flags
   flag.StringVar(&f.repo, "r", "", "repository")
   flag.Func("t", "topic", func(topic string) error {
      f.topics = append(f.topics, topic)
      return nil
   })
   flag.Parse()
   if f.repo != "" {
      res, err := f.put()
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
   } else {
      flag.Usage()
   }
}
