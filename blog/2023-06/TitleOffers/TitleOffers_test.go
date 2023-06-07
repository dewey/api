package justwatch

import (
   "fmt"
   "testing"
)

func Test_Offer(t *testing.T) {
   offers, err := title_offers("tse371404")
   if err != nil {
      t.Fatal(err)
   }
   for offer := range compact(offers) {
      fmt.Print(offer, "\n\n")
   }
}
