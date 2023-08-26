package segments

type segment struct {
	Id   int64  `db:"id"`
	Slug string `db:"slug"`
}
