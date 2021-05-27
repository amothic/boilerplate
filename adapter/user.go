package adapter

import (
	"database/sql"

	"github.com/amothic/boilerplate/repository"

	"github.com/amothic/boilerplate/entity"
)

type userRepoImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepoImpl{
		db: db,
	}
}

func (u *userRepoImpl) GetAll() ([]*entity.User, error) {
	rows, _ := u.db.Query("SELECT name FROM users")
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		users = append(users, entity.NewUser(name))
	}
	return users, nil
}

func (u *userRepoImpl) Create(user *entity.User) error {
	_, err := u.db.Exec("INSERT INTO users(name) VALUES(?)", user.Name())
	return err
}
