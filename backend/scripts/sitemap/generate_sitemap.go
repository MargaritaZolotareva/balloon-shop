package main

import (
	db2 "backend/infrastructure/db"
	"fmt"
	"log"
	"os"
)

func main() {
	db := db2.InitDB()

	sqlDB, err := db.DB()
	if err != nil {
		sqlDB.Close()
	}
	defer sqlDB.Close()

	var sitemap string
	err = sqlDB.QueryRow(`
        SELECT xmlroot(
          xmlelement(
            name urlset,
            xmlattributes('http://www.sitemaps.org/schemas/sitemap/0.9' AS xmlns),
            xmlelement(
              name url,
              xmlelement(name loc, 'https://shariki-v-permi.ru/'),
              xmlelement(name lastmod, current_date),
              xmlelement(name changefreq, 'daily'),
              xmlelement(name priority, '1.0')
            ),
            (
              SELECT xmlagg(
                xmlelement(
                  name url,
                  xmlelement(name loc, 'https://shariki-v-permi.ru/categories/' || slug || '/products'),
                  xmlelement(name lastmod, current_date),
                  xmlelement(name changefreq, 'weekly'),
                  xmlelement(name priority, '0.8')
                )
              )
              FROM categories
              WHERE slug IS NOT NULL
            ),
            (
              SELECT xmlagg(
                xmlelement(
                  name url,
                  xmlelement(name loc, 'https://shariki-v-permi.ru/products/' || slug),
                  xmlelement(name lastmod, current_date),
                  xmlelement(name changefreq, 'weekly'),
                  xmlelement(name priority, '0.6')
                )
              )
              FROM products
              WHERE slug IS NOT NULL
            )
          ),
          version '1.0',
          standalone yes
        )
    `).Scan(&sitemap)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("./../frontend/public/sitemap.xml", []byte(sitemap), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sitemap generated successfully!")
}
