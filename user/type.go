package user

type User struct {
	Id     int64
	Name   string `db:"name"`
	Age    int8   `db:"age"`
	RoleId int    `db:"role_id"`
}