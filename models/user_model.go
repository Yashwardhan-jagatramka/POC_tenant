package models

type User struct {
	UniqueId        int    `json:"id,omitempty" validate:"omitempty"`
	TenantFirstName string `json:"fname,omitempty" validate:"required"`
	TenantLastName  string `json:"lname,omitempty" validate:"required"`
	Country         string `json:"location,omitempty" validate:"required"`
	BusinessDomain  string `json:"domain,omitempty" validate:"required"`
	OfficialEmail   string `json:"email,omitempty" validate:"required,email"`
	OfficialPhone   int    `json:"phone,omitempty" validate:"required" `
}
