package model


type ScheduleRequest struct {
	CenterID      int      `json:"center_id"`
	SessionID     string   `json:"session_id"`
	Beneficiaries []string `json:"beneficiaries"`
	Slot          string   `json:"slot"`
	Dose          int      `json:"dose"`
	Captcha       string   `json:"captcha"`
}