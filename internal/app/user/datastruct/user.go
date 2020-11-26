package datastruct

// User ...
type User struct {
	tableName struct{} `sql:"user"`
	ID        int64    `sql:"id"`
	Name      string   `sql:"name"`
}
