package extra

import (
	"encoding/json"
	"fmt"
)

type Asset struct {
	Business string  `json:"business"`
	Code     string  `json:"code"`
	Type     string  `json:"type"`
	Agency   string  `json:"agency"`
	Location string  `json:"location"`
	City     string  `json:"city"`
	Area     string  `json:"area"`
	Price    string  `json:"price"`
	Numrooms string  `json:"numrooms"`
	Numbaths string  `json:"numbaths"`
	Status   bool    `json:"status"`
	Link     string  `json:"link"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}

func (a *Asset) ToJSON() string {
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
func (a *Asset) FromJSON(s string) {
	input := []byte(s)
	err := json.Unmarshal(input, a)
	if err != nil {
		fmt.Println("error:", err)
	}
}
