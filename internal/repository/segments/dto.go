package segments

type Segment struct {
	Id           int64  `db:"id"`
	Slug         string `db:"slug"`
	PercentUsers int8   `db:"percentage_users"`
}
