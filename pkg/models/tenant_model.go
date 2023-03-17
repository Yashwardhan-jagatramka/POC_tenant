package models

type Tenant struct {
	UniqueId        int    `json:"id" bson:"id" validate:"omitempty"`
	TenantFirstName string `json:"fname" bson:"fname" validate:"required"`
	TenantLastName  string `json:"lname" bson:"lname" validate:"required"`
	Country         string `json:"location" bson:"location" validate:"required"`
	BusinessDomain  string `json:"domain" bson:"domain" validate:"required"`
	OfficialEmail   string `json:"email" bson:"email" validate:"required,email"`
	OfficialPhone   int    `json:"phone" bson:"phone" validate:"required" `
}
