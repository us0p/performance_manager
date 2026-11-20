package entities

type CreateUser struct {
	Email    string
	Name     string
	Password string
}

func NewCreateUserEnt(email, name, pass string) CreateUser {
	return CreateUser{
		email,
		name,
		pass,
	}
}

type User struct {
	Email string
	Name  string
	ID    string
}
