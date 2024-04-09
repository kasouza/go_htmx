package users

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id    int
	Login string
}

func FindByEmail(email string) (*User, error) {
	db, err := sql.Open("mysql", "root:root@/regulatorio")
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT id, login FROM users WHERE email = ?", email)

	var user User
	err = row.Scan(&user.Id, &user.Login)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func All() ([]User, error) {
	db, err := sql.Open("mysql", "root:root@/regulatorio")
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, login FROM users")
	if err != nil {
		return nil, err
	}

	users := []User{}

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Login)

		users = append(users, user)
	}

	return users, nil
}
