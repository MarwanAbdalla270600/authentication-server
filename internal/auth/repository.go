package auth

import "github.com/jmoiron/sqlx"

type RepoInterface interface {
	GetUserByEmail(email string) (*UserDAO, error)
	CreateUser(data *UserDAO) (*UserDAO, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) RepoInterface {
	return &repository{
		db: db,
	}
}

func (r *repository) GetUserByEmail(email string) (*UserDAO, error) {
	var user UserDAO
	err := r.db.Get(&user, "`SELECT * FROM users WHERE email = ?`, email")
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) CreateUser(data *UserDAO) (*UserDAO, error) {
	query := `
        INSERT INTO users (id, first_name, last_name, email, password)
        VALUES (:id, :first_name, :last_name, :email, :password)
    `
	_, err := r.db.NamedExec(query, data)
	if err != nil {
		return nil, err
	}

	// Re-fetch the row to get timestamps
	var user UserDAO
	err = r.db.Get(&user, "SELECT * FROM users WHERE id = ?", data.Id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
