package client

// Credential is common request parameter of Face++ API.
type Credential struct {
	APIKey    string `url:"api_key"`
	APISecret string `url:"api_secret"`
}

// BaseResponse is common response of Face++ API.
type BaseResponse struct {
	RequestID    string `url:"request_id" json:"request_id"`
	TimeUsed     int    `url:"time_used" json:"time_used"`
	ErrorMessage string `url:"error_message" json:"error_message"`
}
