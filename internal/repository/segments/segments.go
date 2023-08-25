package segments

import "github.com/jmoiron/sqlx"

type DbSegments struct {
	conn *sqlx.DB
}

func NewDbSegment(conn *sqlx.DB) *DbSegments {
	return &DbSegments{
		conn: conn,
	}
}

const createSegmentQuery = `INSERT INTO segments(slug) VALUES ($1) ON CONFLICT (slug) DO NOTHING`

func (d *DbSegments) CreateSegment(slug string) error {
	_, err := d.conn.Exec(createSegmentQuery, slug)

	if err != nil {
		return err
	}

	return nil
}
