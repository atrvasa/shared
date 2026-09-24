package dto

// ServiceIdentity نشان‌دهنده هویت یک سرویس داخلی است که با Service-Key احراز شده است
type ServiceIdentity struct {
	ServiceName string `json:"service_name"`
	IsSystem    bool   `json:"is_system"` // برای تشخیص سرویس‌های سیستمی سطح بالا
}

// UserIdentity نشان‌دهنده هویت یک کاربر انسانی است که با JWT یا Session احراز شده است
type UserIdentity struct {
	Username string   `json:"username"`
	FullName string   `json:"full_name"`
	RoleIds  []string `json:"role_ids"`
}
