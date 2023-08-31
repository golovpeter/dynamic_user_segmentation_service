package segments

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/golovpeter/avito-trainee-task-2023/internal/cache/percent_segments"
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
	INSERT INTO segments(slug, percentage_users) 
	VALUES ($1, $2) 
	ON CONFLICT (slug) DO 
	UPDATE SET deleted = false, updated_at = now()
`

func (s *repository) CreateSegment(slug string, percentageUsers int64) error {
	if percentageUsers == 0 {
		_, err := s.conn.Exec(createSegmentQuery, slug, nil)
		return err
	}

	_, err := s.conn.Exec(createSegmentQuery, slug, percentageUsers)
	return err
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

func (s *repository) DeleteSegment(slug string) (bool, error) {
	tx, err := s.conn.Begin()
	if err != nil {
		return false, err
	}

	var slugId int64
	err = tx.QueryRow(deleteSegmentQuery, slug).Scan(&slugId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}

	if slugId == 0 {
		_ = tx.Rollback()
		return false, nil
	}

	_, err = tx.Exec(deleteUserSegmentsQuery, slugId)
	if err != nil {
		_ = tx.Rollback()
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}

const getActiveSegmentsIdsBySlugsQuery = `
	SELECT id, slug
	FROM segments
	WHERE slug IN (?) AND deleted = false
`

func (s *repository) GetActiveSegmentsIdsBySlugs(slugs []string) (map[string]int64, error) {
	var segments []Segment

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

const getPercentSegmentsQuery = `
	SELECT id, slug, percentage_users
	FROM segments
	WHERE percentage_users IS NOT NULL AND deleted = false 
`

func (s *repository) GetPercentSegments() (map[string]percent_segments.Segment, error) {
	var segments []Segment

	err := s.conn.Select(&segments, getPercentSegmentsQuery)
	if err != nil {
		return nil, err
	}

	percentSegments := make(map[string]percent_segments.Segment)

	for _, val := range segments {
		percentSegments[val.Slug] = percent_segments.Segment{
			Id:           val.Id,
			Slug:         val.Slug,
			PercentUsers: val.PercentUsers,
		}
	}

	return percentSegments, err
}
