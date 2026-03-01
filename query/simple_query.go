package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	ErrDb = errors.New("error db")
)

type FuckDB struct {
}

func (f FuckDB) Error() string {
	return "fuck"
}

func main() {
	partners, err := GetPartners()
	if err != nil {
		if errors.Is(err, ErrDb) {
			fmt.Println("fuck u")
			return
		}
		log.Fatal(err)
	}
	fmt.Println(partners)
}

func GetPartners() ([][]interface{}, error) {
	fuckDb := FuckDB{}
	dsn := "root:root@tcp(127.0.0.1:3307)/b2b"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		err := db.Close()
		log.Println(err)
	}()

	query := `
SELECT
    partner.id partner_id,
    partner.title partner_title,
    sale_point.id sale_point_id,
    address.address store_address,
    CASE
        WHEN partner_disable_store.id IS NULL THEN 'Да'
        ELSE 'Нет'
    END enable_store
FROM partner
CROSS JOIN sale_point
LEFT JOIN partner_disable_store
       ON partner_disable_store.idPartner = partner.id
      AND partner_disable_store.idSalePoint = sale_point.id
LEFT JOIN address
       ON address.id = sale_point.idAddress
WHERE partner.deleted = 0
AND sale_point.isActive = 1
ORDER BY partner.id, sale_point.id;
`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("error while sql get partner : %w")
		return nil, fmt.Errorf("error get partner: %w", fuckDb)
	}
	defer func() {
		err := rows.Close()
		log.Println(err)
	}()

	cols, err := rows.Columns()
	if err != nil {
		log.Println(err)
	}

	result := make([]interface{}, len(cols))
	links := make([]interface{}, len(cols))

	queryResult := make([][]interface{}, len(cols))

	for i := range result {
		links[i] = &result[i]
	}

	for rows.Next() {
		if err := rows.Scan(links...); err != nil {
			fmt.Println(err)
		}
		fmt.Println(result...)
		queryResult = append(queryResult, result)
	}
	return queryResult, nil
}
