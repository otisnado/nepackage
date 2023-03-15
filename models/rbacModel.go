package models

type Role struct {
	ID   uint   `json:"id" gorm:"primary_key,not null,auto_increment"`
	Name string `json:"name" binding:"required" gorm:"not null"`
}

type RoleUpdate struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty" binding:"required"`
}

type SuccessFindRoles struct {
	Data []Role `json:"data"`
}

type SuccessRoleCreation struct {
	Data Role `json:"data"`
}

type SuccessRoleUpdate struct {
	Data Role `json:"data"`
}

type SuccessRoleDelete struct {
	Data bool `json:"data"`
}

type Policy struct {
	ID                uint   `json:"id" gorm:"primary_key,not null,auto_increment"`
	Name              string `json:"name" binding:"required" gorm:"not null"`
	Path              string `json:"path" binding:"required" gorm:"not null"`
	AuthorizedMethods string `json:"authorizedMethods" binding:"required" gorm:"not null"`
	AuthorizedRole    uint   `json:"authorizedRole" binding:"required" gorm:"not null"`
}

type PolicyUpdate struct {
	ID             uint   `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Path           string `json:"path,omitempty"`
	AuthorizedRole uint   `json:"authorizedRole,omitempty"`
}

type SuccessFindPolicies struct {
	Data []Policy `json:"data"`
}

type SuccessPolicyCreation struct {
	Data Policy `json:"data"`
}

type SuccessPolicyUpdate struct {
	Data Policy `json:"data"`
}

type SuccessPolicyDelete struct {
	Data bool `json:"data"`
}
