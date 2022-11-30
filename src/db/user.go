package db

const USER_TABLE_NAME = "user"

type UserRoleType int32

const (
	ADMIN UserRoleType = iota + 1
	COMMON
	RESTRICTED
)

const (
	ADMIN_STR = "管理员"

	COMMON_STR = "操作员"

	RESTRICTED_STR = "访客"
)

var UserRoleTypeName = map[UserRoleType]string{
	ADMIN:      ADMIN_STR,
	COMMON:     COMMON_STR,
	RESTRICTED: RESTRICTED_STR,
}

var UserRoleTypeValue = map[string]UserRoleType{
	ADMIN_STR:      ADMIN,
	COMMON_STR:     COMMON,
	RESTRICTED_STR: RESTRICTED,
}

type User struct {
	GeneralField
	UserName     string `gorm:"index"`
	UserRole     UserRoleType
	UserPwd      string
	UserNickName string
	UserPhoneNum string
	UserEmail    string
}

func (u *User) TableName() string {
	return USER_TABLE_NAME
}

func init() {
	user := new(User)
	TableSlice = append(TableSlice, &user)
}
