package segments

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DbSegments struct {
	conn *sqlx.DB
}

func NewDbSegment(conn *sqlx.DB) *DbSegments {
	return &DbSegments{
		conn: conn,
	}
}

const createSegmentQuery = `INSERT INTO segments(slug) VALUES ($1) 
                           ON CONFLICT (slug) DO UPDATE SET deleted = false, updated_at = now()`

func (d *DbSegments) CreateSegment(slug string) error {
	_, err := d.conn.Exec(createSegmentQuery, slug)
	if err != nil {
		return err
	}

	return nil
}

const deleteSegmentQuery = `UPDATE segments SET deleted = true, updated_at = now() 
                WHERE slug = $1`

func (d *DbSegments) DeleteSegment(slug string) (error, sql.Result) {
	result, err := d.conn.Exec(deleteSegmentQuery, slug)
	if err != nil {
		return err, nil
	}

	return nil, result
}
