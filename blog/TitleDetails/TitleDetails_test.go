package justwatch

import (
   "fmt"
   "testing"
)

func Test_Detail(t *testing.T) {
   //full_path := "/us/tv-show/orphan-black/season-1"
   full_path := "/us/tv-show/orphan-black"
   detail, err := new_title_details(full_path)
   if err != nil {
      t.Fatal(err)
   }
   for _, episode := range detail.Data.URL.Node.Episodes {
      fmt.Printf("%+v\n", episode)
   }
}
