package userServices

import "database/sql"

type User struct {
	Id     int64
	Name   string `db:"name"`
	Age    int8   `db:"age"`
	RoleId int    `db:"role_id"`
}

func Insert(u *User, db *sql.DB) (int64, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	result, err := db.Exec("insert into user (name, age, role_id) values (?, ?, ?)", u.Name, u.Age, u.RoleId)
	if err != nil {
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	u.Id = affected
	return affected, nil
}
