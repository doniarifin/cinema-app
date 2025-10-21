package model

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
	GuestRole Role = "guest"
)

func (r Role) String() string {
	return string(r)
}
