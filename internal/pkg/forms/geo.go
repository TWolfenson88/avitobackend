package forms

type GeoIP struct {
	Ip          string  `json:"ip"`
	Country		string  `json:"country"`
	Region      string	`json:"region"`
	City		string	`json:"city"`
}
