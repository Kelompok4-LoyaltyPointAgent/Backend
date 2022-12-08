package constant

type UserRoleEnum string

const (
	UserRoleUser  UserRoleEnum = "User"
	UserRoleAdmin UserRoleEnum = "Admin"
)

func (x UserRoleEnum) String() string {
	return string(x)
}
