package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
)

type Row struct {
	PartnerID        int
	PartnerTitle     string
	SalePointID      int
	SalePointAddress sql.NullString
	AllowDelivery    string
}

type StoreInfo struct {
	SalePointID    int
	SalePointLabel string
	AllowDelivery  string
}

type PartnerDelivery struct {
	PartnerID    int
	PartnerLabel string
	Stores       map[int]StoreInfo // key = sale_point_id
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3307)/b2b?charset=utf8mb4&parseTime=true&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("sql.Open error: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("db.Ping error: %v", err)
	}

	query := `
SELECT
   partner.id partner_id,
   partner.title partner_title,
   sale_point.id sale_point_id,
   address.address sale_point_address,
   CASE
       WHEN partner_disable_store.id IS NULL THEN 'Да'
       ELSE 'Нет'
   END allow_delivery
FROM partner
CROSS JOIN sale_point
LEFT JOIN partner_disable_store
      ON partner_disable_store.idPartner = partner.id
     AND partner_disable_store.idSalePoint = sale_point.id
LEFT JOIN address
      ON address.id = sale_point.idAddress
WHERE partner.deleted = 0
ORDER BY partner.id, sale_point.id;
`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("db.Query error: %v", err)
	}

	fmt.Println(rows)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	t := reflect.TypeOf(rows)
	// Если это указатель, получаем тип, на который он указывает
	if t.Kind() == reflect.Ptr {
		t = t.Elem() // Разыменовываем указатель
	}

	fmt.Printf("%T\n", rows)

}
