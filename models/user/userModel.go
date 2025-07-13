package userModel

import "github.com/githubak2002/golang-event-api/db"

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}

	defer stmt.Close()

	res,err := stmt.Exec(user.Email, user.Password)
	if err != nil{
		return err
	}

	userId,err := res.LastInsertId()
	user.Id = userId

	return err
}