package models

type OtpData struct {
	Otp   string `json:"otp"`
	Gmail string `json:"gmail"`
}

type CheckOTPResp struct {
	IsRight bool `json:"is_right"`
}
