package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/meez25/boilerplateForumDDD/internal/authentication"
)

type AuthenticationSQLRepository struct {
	conn *SQLConnection
}

func NewAuthenticationSQLRepository(conn *SQLConnection) *AuthenticationSQLRepository {

	err := conn.conn.Ping(context.Background())

	if err != nil {
		fmt.Println(err)
	}

	_, err = conn.conn.Exec(context.Background(), `
    CREATE TABLE IF NOT EXISTS sessions
        (
            id UUID PRIMARY KEY,
            email VARCHAR(100) NOT NULL,
            userid VARCHAR(100) NOT NULL,
            username VARCHAR(50) NOT NULL,
            valid_until TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        );
    `)

	if err != nil {
		fmt.Println(err)
	}

	return &AuthenticationSQLRepository{
		conn: conn,
	}
}

func (ar *AuthenticationSQLRepository) Save(session authentication.Session) error {
	_, err := ar.conn.conn.Exec(context.Background(), `
        INSERT INTO sessions VALUES ($1, $2, $3, $4, $5)
`, session.ID, session.Email, session.UserID, session.Username, session.GetValidUntil())

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				fmt.Println(pgErr.Message)
				// switch pgErr.ConstraintName {
				// case "users_username_key":
				// 	return ErrDuplicateUsername
				// case "users_email_key":
				// 	return ErrDuplicateEmail
				// default:
				// 	return ErrUserAlreadyExists
				// }
			}
		}
		return err
	}
	return nil
}

func (ar *AuthenticationSQLRepository) FindByID(ID string) (authentication.Session, error) {
	var s authentication.Session

	row := ar.conn.conn.QueryRow(context.Background(), `
        SELECT * from sessions WHERE id = $1
        `,
		ID)
	row.Scan(&s.ID, &s.Email, &s.UserID, &s.Username, &s.ValidUntil)

	if s.ID == uuid.Nil {
		return authentication.Session{}, fmt.Errorf("Not found")
	}

	return s, nil
}

func (ar *AuthenticationSQLRepository) Update(authentication.Session) error {
	return nil
}

func (ar *AuthenticationSQLRepository) Delete(ID string) error {
	return nil
}
