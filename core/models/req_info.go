package models

type ReqInfo struct {
	Token        string `json:"token"`
	Username     string `json:"username"`
	IsActive     bool   `json:"is_active"`
	IsSuperAdmin bool   `json:"is_super_admin"`
	User         *User  `json:"user"`
}

func NewReqInfo(token string, user *User) *ReqInfo {
	return &ReqInfo{
		Token:        token,
		Username:     user.Username,
		IsActive:     user.IsActive,
		IsSuperAdmin: user.Role == UserRoleSuperAdmin,
		User:         user,
	}
}
