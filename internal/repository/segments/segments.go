package segments

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{
		conn: conn,
	}
}

const createSegmentQuery = `
	INSERT INTO segments(slug) 
	VALUES ($1) 
	ON CONFLICT (slug) DO 
	UPDATE SET deleted = false, updated_at = now()
`

func (s *repository) CreateSegment(slug string) error {
	_, err := s.conn.Exec(createSegmentQuery, slug)
	if err != nil {
		return err
	}

	return nil
}

const deleteSegmentQuery = `
	UPDATE segments SET deleted = true, updated_at = now() 
    WHERE slug = $1
`

func (s *repository) DeleteSegment(slug string) (int64, error) {
	result, err := s.conn.Exec(deleteSegmentQuery, slug)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

const getActiveSegmentsIdsBySlugsQuery = `
	SELECT id, slug
	FROM segments
	WHERE slug IN (?) AND deleted = false
`

func (s *repository) GetActiveSegmentsIdsBySlugs(slugs []string) (map[string]int64, error) {
	var segments []segment

	query, args, err := sqlx.In(getActiveSegmentsIdsBySlugsQuery, slugs)
	if err != nil {
		return nil, err
	}

	query = s.conn.Rebind(query)

	err = s.conn.Select(&segments, query, args...)
	if err != nil {
		return nil, err
	}

	slugsWithIds := make(map[string]int64, len(slugs))

	for _, seg := range segments {
		slugsWithIds[seg.Slug] = seg.Id
	}

	return slugsWithIds, nil
}
