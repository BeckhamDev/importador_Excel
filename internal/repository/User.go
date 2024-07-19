package repository

import (
	"database/sql"
	"importador_Excel/internal/models"
)

//Instancia de conexão com o banco
type User struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

//CreateUser recebe os dados de um novo usuário e salva no banco retornando o usuario
func (db *User) CreateUser(User models.User) (models.User, error) {
	sql := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	stmt, err := db.db.Prepare(sql)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(User.Name, User.Email, User.Password)
	if err != nil {
		return models.User{}, err
	}

	return User, err
}

//GetUserByEmail busca um usuário no banco a partir de seu email
func (db *User) GetUserByEmail(email string) (models.User, error) {
	sql := "SELECT * FROM users WHERE email = $1"
	stmt, err := db.db.Query(sql, email)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	var user models.User

	if stmt.Next(){
		if err := stmt.Scan(
			&user.ID, 
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}