package domain

// ユーザ
type User struct {
	*Model
	Name  string
	Email string
	Uid   string
}
