package repository

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type UserActivity struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	DateStart   time.Time `json:"date_start"`
	DateEnd     time.Time `json:"date_end"`
	ProjectID   int32     `json:"project_id"`
	ProjectName string    `json:"project_name"`
	UserID      int32     `json:"user_id"`
}

func (q *Queries) GetUserActivites(ctx context.Context, projects []string, search string, userId int32) ([]UserActivity, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.Select("a.id, a.name, a.date_end, a.date_start, p.id, p.name, a.user_id").From("activities AS a").Join("projects AS p on a.project_id = p.id").Where(sq.Eq{"a.user_id": userId})
	if len(projects) > 0 {
		placeholders := make([]interface{}, len(projects))
		for i, project := range projects {
			placeholders[i] = project
		}

		query = query.Where(sq.Eq{"p.name": placeholders})
	}
	if search != "" {
		query = query.Where(sq.Expr("a.name ILIKE ?", fmt.Sprintf("%%%s%%", search)))
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := q.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserActivity{}
	for rows.Next() {
		var i UserActivity
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.DateEnd,
			&i.DateStart,
			&i.ProjectID,
			&i.ProjectName,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil

}
