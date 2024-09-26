package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/meez25/boilerplateForumDDD/internal/user"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrDuplicateUsername = errors.New("username already exists")
	ErrDuplicateEmail    = errors.New("email already exists")
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
`, user.ID, user.Username, user.EmailAddress, user.Password.Password, user.FirstName, user.LastName, user.ProfilePic, user.CreatedAt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				switch pgErr.ConstraintName {
				case "users_username_key":
					return ErrDuplicateUsername
				case "users_email_key":
					return ErrDuplicateEmail
				default:
					return ErrUserAlreadyExists
				}
			}
		}
		return err
	}

	return nil
}

func (ur *UserSQLRepository) FindByID(ID string) (user.User, error) {
	return user.User{}, nil
}

func (ur *UserSQLRepository) FindByEmailAddress(email string) (user.User, error) {
	var u user.User

	row := ur.conn.conn.QueryRow(context.Background(), `
        SELECT * from users WHERE email = $1
        `,
		email)
	row.Scan(&u.ID, &u.Username, &u.EmailAddress, &u.Password.Password, &u.FirstName, &u.LastName, &u.ProfilePic, &u.CreatedAt)

	if u.ID == uuid.Nil {
		return user.User{}, fmt.Errorf("Not found")
	}

	return u, nil
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
