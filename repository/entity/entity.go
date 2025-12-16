package entityrepository

import (
	"context"
	"fmt"
	"ostadbun/interface/entity"

	"github.com/jackc/pgx/v5"
)

type EntityRepo struct {
	db *pgx.Conn
}

func Make(db *pgx.Conn) EntityRepo {
	return EntityRepo{
		db: db,
	}

}

func (r EntityRepo) Search(query string) ([]entity.IEntity, error) {
	ctx := context.Background()
	var results []entity.IEntity

	// جداولی که name مستقیم دارند
	tables := []struct {
		name     string
		category string
	}{
		{"professor", "professor"},
		{"major", "major"},
		{"university", "university"},
		{"lesson", "lesson"},
	}

	for _, t := range tables {
		sql := fmt.Sprintf("SELECT id, name FROM %s WHERE name ILIKE $1", t.name)
		rows, err := r.db.Query(ctx, sql, "%"+query+"%")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				return nil, err
			}
			results = append(results, entity.IEntity{
				Name:     name,
				Category: t.category,
			})
		}
	}

	// جداولی که name داخل JSONB wiki هستند (مثلاً lesson)
	sql := "SELECT id, wiki->>'name' AS name FROM lesson WHERE wiki->>'name' ILIKE $1"
	rows, err := r.db.Query(ctx, sql, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		results = append(results, entity.IEntity{
			Name:     name,
			Category: "lesson",
		})
	}

	return results, nil
}
