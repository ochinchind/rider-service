package ride

import (
	"encoding/json"
	"time"
)

type Location struct {
	CreatedAt time.Time `json:"time" redis:"time"`
	Latitude  float64   `json:"latitude" redis:"latitude"`
	Longitude float64   `json:"longitude" redis:"longitude"`
}

func (l *Location) MarshalBinary() ([]byte, error) {
	return json.Marshal(l)
}

func (l *Location) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, l)
}
