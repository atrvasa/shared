package dto

// PolicyCheckRequest درخواستی برای بررسی دسترسی
type PolicyCheckRequest struct {
	SubjectID   uint   `json:"subject_id"`   // UserID یا ServiceID
	SubjectType string `json:"subject_type"` // "user" or "service"
	Resource    string `json:"resource"`     // مثلا: "sms_service"
	Action      string `json:"action"`       // مثلا: "send_otp"
}

// PolicyCheckResponse پاسخ موتور قوانین
type PolicyCheckResponse struct {
	Allowed bool   `json:"allowed"`
	Reason  string `json:"reason,omitempty"`
}
