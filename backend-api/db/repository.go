package db

import (
	// ...
	"fmt"
	"vue-golang-payment-app/backend-api/domain"
)

// Select all Items
func SelectAllItems() (items domain.Items, err error) {
	stmt, err := Conn.Query("SELECT * FROM items")
	if err != nil {
		return
	}
	defer stmt.Close()
	for stmt.Next() {
		var id int64
		var name string
		var description string
		var amount int64
		if err := stmt.Scan(&id, &name, &description, &amount); err != nil {
			continue
		}
		item := domain.Item{
			ID:          id,
			Name:        name,
			Description: description,
			Amount:      amount,
		}
		items = append(items, item)
	}
	return
}

// 以下のバグを直す
func SelectItem(identifier int64) (item domain.Item, err error) {
	stmt, err := Conn.Query("SELECT * FROM items WHERE id = ? LIMIT 1", identifier)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	var id int64
	var name string
	var description string
	var amount int64
	// ポインタ渡しにして、値の変更を保持するために&をつける。ここと上のQueryがダメ
	if err = stmt.Scan(&id, &name, &description, &amount); err != nil {
		return
	}
	item.ID = id
	item.Name = name
	item.Description = description
	item.Amount = amount
	return
}
