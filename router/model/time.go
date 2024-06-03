package model

import (
	"encoding/json"
	"log"
	"time"
)

// TODO: check if this is required or not
type customTime time.Time // created since time.Time does not bind properly for gin

func (t *customTime) UnmarshalJSON(bs []byte) error {
	var timestamp int64
	err := json.Unmarshal(bs, &timestamp)
	if err != nil {
		return err
	}

	*t = customTime(time.Unix(timestamp/1000, timestamp%1000*1e6))
	return nil
}

func (t customTime) MarshalJSON() ([]byte, error) {
	timestamp := time.Time(t).UnixNano() / 1e6
	log.Println(time.Time(t).UnixNano())
	return json.Marshal(timestamp)
}
