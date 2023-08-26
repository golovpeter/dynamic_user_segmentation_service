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
			Columns("user_id", "segment_id").PlaceholderFormat(squirrel.Dollar)

		for _, segmentId := range changeData.AddSegmentsIds {
			insertBuilder = insertBuilder.Values(changeData.UserID, segmentId)
		}

		insertBuilder = insertBuilder.Suffix("ON CONFLICT DO NOTHING")

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
