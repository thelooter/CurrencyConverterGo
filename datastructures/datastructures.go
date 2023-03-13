package datastructures

type Config struct {
	APIKey  string `json.json:"APIKey"`
	BaseURL string `json.json:"baseURL"`
}

type Request struct {
	Result             string  `json:"result"`
	Documentation      string  `json:"documentation"`
	TermsOfUse         string  `json:"terms_of_use"`
	TimeLastUpdateUnix int     `json:"time_last_update_unix"`
	TimeLastUpdateUtc  string  `json:"time_last_update_utc"`
	TimeNextUpdateUnix int     `json:"time_next_update_unix"`
	TimeNextUpdateUtc  string  `json:"time_next_update_utc"`
	BaseCode           string  `json:"base_code"`
	TargetCode         string  `json:"target_code"`
	ConversionRate     float64 `json:"conversion_rate"`
}
