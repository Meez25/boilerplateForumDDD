package persistence

import (
	"context"
	"fmt"

	"github.com/meez25/boilerplateForumDDD/internal/user"
)

type UserSQLRepository struct {
	conn *SQLConnection
}

func NewUserSQLRepository(conn *SQLConnection) *UserSQLRepository {

	err := conn.conn.Ping(context.Background())

	if err != nil {
		fmt.Println(err)
	}

	_, err = conn.conn.Exec(context.Background(), `
    CREATE TABLE IF NOT EXISTS users 
        (
            id UUID PRIMARY KEY,
            username VARCHAR(50) UNIQUE NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            password VARCHAR(100) NOT NULL,
            first_name VARCHAR(50),
            last_name VARCHAR(50),
            profile_pic VARCHAR(255),
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        );
    `)

	if err != nil {
		fmt.Println(err)
	}

	return &UserSQLRepository{
		conn: conn,
	}
}

func (ur *UserSQLRepository) Save(user user.User) error {
	_, err := ur.conn.conn.Exec(context.Background(), `
        INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`, user.ID, user.Username, user.EmailAddress, user.Password, user.FirstName, user.LastName, user.ProfilePic, user.CreatedAt)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (ur *UserSQLRepository) FindByID(ID string) (user.User, error) {
	return user.User{}, nil
}

func (ur *UserSQLRepository) FindByEmailAddress(email string) (user.User, error) {
	var user user.User

	row := ur.conn.conn.QueryRow(context.Background(), `
        SELECT * from users WHERE email = $1
        `,
		email)
	row.Scan(&user.ID, &user.Username, &user.EmailAddress, &user.Password, &user.FirstName, &user.LastName, &user.ProfilePic, &user.CreatedAt)
	return user, nil
}

func (ur *UserSQLRepository) FindByUsername(username string) (user.User, error) {
	return user.User{}, nil
}

func (ur *UserSQLRepository) Update(user user.User) error {
	return nil
}

func (ur *UserSQLRepository) Delete(ID string) error {
	return nil
}
