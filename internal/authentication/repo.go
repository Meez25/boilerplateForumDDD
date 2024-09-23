package authentication

type SessionRepository interface {
	Save(session Session)
	FindByID(ID string) (Session, error)
	Update(Session Session) error
	Delete(ID string) error
}
