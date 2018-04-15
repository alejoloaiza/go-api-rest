package extra

import (
	"encoding/json"
	"fmt"
)

type Asset struct {
	Business string  `json:"business,omitempty"`
	Code     string  `json:"code,omitempty"`
	Type     string  `json:"type,omitempty"`
	Agency   string  `json:"agency,omitempty"`
	Location string  `json:"location,omitempty"`
	City     string  `json:"city,omitempty"`
	Area     string  `json:"area,omitempty"`
	Price    string  `json:"price,omitempty"`
	Numrooms string  `json:"numrooms,omitempty"`
	Numbaths string  `json:"numbaths,omitempty"`
	Status   bool    `json:"status,omitempty"`
	Link     string  `json:"link,omitempty"`
	Lat      float64 `json:"lat,omitempty"`
	Lon      float64 `json:"lon,omitempty"`
}

func (a *Asset) ToJSON() string {
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
func (a *Asset) FromJSON() string {
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
