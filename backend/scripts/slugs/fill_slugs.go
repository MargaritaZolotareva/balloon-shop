package main

import (
	db2 "backend/infrastructure/db"
	"database/sql"
	"fmt"

	"github.com/gosimple/slug"
	_ "github.com/lib/pq"
)

// пример запуска: POSTGRES_USER=user POSTGRES_PASSWORD=pswd DB_HOST=host DB_PORT=5432 POSTGRES_DB=db go run scripts/slugs/fill_slugs.go

func main() {
	db := db2.InitDB()

	sqlDB, err := db.DB()
	if err != nil {
		sqlDB.Close()
	}
	defer sqlDB.Close()
	updateSlugs(sqlDB, "categories")
	updateSlugs(sqlDB, "products")

	fmt.Println("All slugs updated!")
}

func updateSlugs(sqlDB *sql.DB, table string) {
	rows, err := sqlDB.Query("SELECT id, title FROM " + table)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	type pair struct {
		id   int
		slug string
	}
	var updates []pair
	for rows.Next() {
		var id int
		var title string
		if err := rows.Scan(&id, &title); err != nil {
			panic(err)
		}
		slugValue := fmt.Sprintf("%s-%d", slug.Make(title), id)
		updates = append(updates, pair{id, slugValue})
	}
	if len(updates) == 0 {
		return
	}

	values := ""
	args := []interface{}{}

	for i, u := range updates {
		if i > 0 {
			values += ","
		}
		values += fmt.Sprintf("($%d::int, $%d::text)", i*2+1, i*2+2)
		args = append(args, u.id, u.slug)
	}

	query := fmt.Sprintf(
		"UPDATE %s t "+
			"SET slug = v.slug "+
			"FROM (VALUES %s) AS v(id, slug) "+
			"WHERE t.id = v.id", table, values)

	_, err = sqlDB.Exec(query, args...)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated %d rows\n", len(updates))

	// делаем другой slug для раздела Акции
	sqlDB.Exec("UPDATE "+table+" SET slug = $1 WHERE id = $2", "promo", 0)
}
