package sqlxmodels

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
	Posts []Post `db:"-"`
}
