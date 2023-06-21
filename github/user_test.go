package github

import (
   "fmt"
   "testing"
)

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
