package user_segments

import (
	"github.com/Masterminds/squirrel"
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

const deleteSegmentsQuery = `
	DELETE FROM users_to_segments 
    WHERE user_id = ? AND segment_id IN (?);
`

func (u *repository) ChangeUserSegments(changeData ChangeUserSegmentsData) error {
	tx, err := u.conn.Begin()
	if err != nil {
		return err
	}

	if len(changeData.AddSegmentsIds) != 0 {
		insertBuilder := squirrel.Insert("users_to_segments").
			Columns("user_id", "segment_id", "expired_at", "added_to_segment").PlaceholderFormat(squirrel.Dollar)

		for _, segmentId := range changeData.AddSegmentsIds {
			if changeData.ExpiredAt.IsZero() {
				insertBuilder = insertBuilder.Values(changeData.UserID, segmentId, nil, true)
				continue
			}

			insertBuilder = insertBuilder.Values(changeData.UserID, segmentId, changeData.ExpiredAt, true)
		}

		if changeData.ExpiredAt.IsZero() {
			insertBuilder = insertBuilder.Suffix("ON CONFLICT (user_id, segment_id) DO UPDATE SET expired_at = null")
		} else {
			insertBuilder = insertBuilder.Suffix("ON CONFLICT (user_id, segment_id) DO UPDATE SET expired_at = ?", changeData.ExpiredAt)
		}

		query, args, err := insertBuilder.ToSql()
		if err != nil {
			return err
		}

		_, err = tx.Exec(query, args...)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if len(changeData.DeleteSegmentsIds) > 0 {
		query, args, err := sqlx.In(deleteSegmentsQuery, changeData.UserID, changeData.DeleteSegmentsIds)
		if err != nil {
			return err
		}

		query = u.conn.Rebind(query)

		_, err = tx.Exec(query, args...)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

const deleteExpiredUserSegmentsQuery = `
	DELETE FROM users_to_segments
	WHERE expired_at <= now();
`

func (u *repository) DeleteExpiredUserSegments() error {
	_, err := u.conn.Exec(deleteExpiredUserSegmentsQuery)
	return err
}

const getUserSegmentsQuery = `
	SELECT segment_id, slug, added_to_segment
	FROM users_to_segments us
	INNER JOIN segments s ON s.id = us.segment_id
	WHERE user_id = $1 
`

func (s *repository) GetUserSegments(id int64) (map[string]SegmentInfo, error) {
	segments := make([]SegmentInfo, 0)
	err := s.conn.Select(&segments, getUserSegmentsQuery, id)

	segmentsMap := make(map[string]SegmentInfo)
	for _, val := range segments {
		segmentsMap[val.Slug] = val
	}

	return segmentsMap, err
}

const addOneUserInSegmentQuery = `
	INSERT INTO users_to_segments (user_id, segment_id, added_to_segment)
	VALUES ($1, $2, $3) 
	ON CONFLICT DO NOTHING 
`

func (s *repository) AddOneUserSegment(userId, segmentId int64, addedSegment bool) error {
	_, err := s.conn.Exec(addOneUserInSegmentQuery, userId, segmentId, addedSegment)

	return err
}
