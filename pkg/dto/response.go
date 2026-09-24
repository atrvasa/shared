package dto

// APIResponse ساختار واحد برای تمام پاسخ‌های HTTP در کل سیستم
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// NewSuccessResponse ایجاد سریع یک پاسخ موفق
func NewSuccessResponse(data interface{}, message string) APIResponse {
	return APIResponse{
		Success: true,
		Data:    data,
		Message: message,
	}
}

// NewErrorResponse ایجاد سریع یک پاسخ خطا
func NewErrorResponse(err string, message string) APIResponse {
	return APIResponse{
		Success: false,
		Error:   err,
		Message: message,
	}
}
