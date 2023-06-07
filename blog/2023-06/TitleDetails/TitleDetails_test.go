package justwatch

import (
   "encoding/json"
   "os"
   "testing"
)

func Test_Detail(t *testing.T) {
   full_path := "/us/tv-show/orphan-black"
   detail, err := new_title_details(full_path)
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   for _, episode := range detail.Data.URL.Node.Episodes {
      enc.Encode(episode)
   }
}
