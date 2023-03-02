package justwatch

import (
   "fmt"
   "testing"
   "time"
)

const enemy = "/us/movie/ennemi"

func Test_Content(t *testing.T) {
   content, err := New_Content_URLs(enemy)
   if err != nil {
     t.Fatal(err)
   }
   offer := make(Offers)
   Client.Log_Level = 2
   for _, tag := range content.Href_Lang_Tags {
      if tag.Language() == "en" {
         v := tag.Variables()
         detail, err := v.Details()
         if err != nil {
            t.Fatal(err)
         }
         offer.Add(v.Country_Code, detail)
         time.Sleep(99 * time.Millisecond)
      }
   }
   text, err := offer.Stream().Text()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(text)
}
