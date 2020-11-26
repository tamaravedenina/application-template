package datastruct

// Book ...
type Book struct {
	tableName struct{} `sql:"book"`
	ID        int64    `sql:"id"`
	UserID    int64    `sql:"user_id"`
	Name      string   `sql:"name"`
}

// BookWithUser ...
type BookWithUser struct {
	ID       int64
	Name     string
	UserName string
}
