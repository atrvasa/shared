package dto

// SMSParameter پارامترهای کلیدی-مقداری برای الگوهای پیامک
type SMSParameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SMSVerifyRequest درخواستی که سرویس‌ها برای ارسال پیامک (معمولاً OTP) ارسال می‌کنند
type SMSVerifyRequest struct {
	PhoneNumber string         `json:"phone_number" binding:"required"`
	TemplateID  string         `json:"template_id" binding:"required"`
	Parameters  []SMSParameter `json:"parameters"`         // لیست پارامترها (مثل کد تایید)
	Provider    string         `json:"provider,omitempty"` // اختیاری: انتخاب ارائه‌دهنده خاص
}

// SMSResponse ساختار پاسخ استاندارد سرویس SMS
type SMSResponse struct {
	MessageID  string `json:"message_id"`
	Status     string `json:"status"` // "queued", "sent", "failed"
	ProviderID string `json:"provider_id,omitempty"`
	Error      string `json:"error,omitempty"`
}
