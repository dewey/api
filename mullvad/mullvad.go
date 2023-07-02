package mullvad

import (
   "encoding/json"
   "io"
)

type server []any

func servers(r io.Reader) ([]server, error) {
   var data struct {
      Data []any
   }
   err := json.NewDecoder(r).Decode(&data)
   if err != nil {
      return nil, err
   }
   var servs []server
   indexes := data.Data[0].([]any)
   for i, low := range indexes {
      low := int(low.(float64))
      high := len(data.Data)-1
      if i+1 < len(indexes) {
         high = int(indexes[i+1].(float64))
      }
      servs = append(servs, data.Data[low:high])
   }
   return servs, nil
}
