package userModel

import (
	"errors"

	"github.com/githubak2002/golang-event-api/db"
	"github.com/githubak2002/golang-event-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`  
	// Password string `binding:"required" json:"-"`  
	// <- will not appear in JSON`
}

func (user *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	res,err := stmt.Exec(user.Email, hashedPassword)
	if err != nil{
		return err
	}

	userId,err := res.LastInsertId()
	user.Id = userId

	return err
}

func GetUsers() ([]User, error) {
	query := `SELECT id, email FROM users`
	rows, err := db.DB.Query(query)
	if err != nil{
		return nil, err
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id,&user.Email)
		if err != nil{
			return nil, err
		}
		users = append(users,user)
	}
	return users,nil
}

func (user *User) ValidateCreadentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)

	// NOTE: QueryRow fetches a single row matching the email.
	// NOTE: Scan reads column values into variables in the same order as the SELECT.

	// Retrieve the hashed password from the database for the given email

	var retrievedPassword string
	err := row.Scan(&user.Id, &retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	isPasswordValid := utils.CheckHashPassword(user.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	return nil	
}