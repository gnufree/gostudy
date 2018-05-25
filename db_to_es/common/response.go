package common


type Response struct {
	Ok bool `json:"ok"`
	ErrorMsg string `json:"errorMsg"`
}