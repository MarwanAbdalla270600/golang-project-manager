package user

type UserRole string

const (
	RoleAdmin   UserRole = "ADMIN"
	RoleManager UserRole = "MANANGER"
	RoleMember  UserRole = "MEMBER"
)

type UserOutput struct {
	Username  string   `json:"username"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Role      UserRole `json:"role"`
}

type LoginUser struct {
	ID       uint     `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

type UserInput struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}
