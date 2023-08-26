package segments

import (
	"database/sql"
	"errors"
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
    WHERE slug = $1 AND deleted = false
	RETURNING id
`

const deleteUserSegmentsQuery = `
	DELETE FROM users_to_segments
	WHERE segment_id = $1
`

func (s *repository) DeleteSegment(slug string) (int64, error) {
	tx, err := s.conn.Begin()
	if err != nil {
		return 0, err
	}

	var slugId int64
	err = tx.QueryRow(deleteSegmentQuery, slug).Scan(&slugId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	if slugId == 0 {
		_ = tx.Rollback()
		return 0, nil
	}

	_, err = tx.Exec(deleteUserSegmentsQuery, slugId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return 1, nil
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

const getUserSegmentsQuery = `
	SELECT slug
	FROM users_to_segments us
	INNER JOIN segments s ON s.id = us.segment_id
	WHERE user_id = $1
`

func (s *repository) GetUserSegments(id int64) ([]string, error) {
	segments := make([]string, 0)

	err := s.conn.Select(&segments, getUserSegmentsQuery, id)
	if err != nil {
		return []string{}, err
	}

	return segments, nil
}
