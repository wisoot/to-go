package item

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"strconv"
)

const itemTable = "items"

type MysqlGateway struct {
	db *sql.DB
}

func (gateway *MysqlGateway) Create(item Item) (Item, error) {
	//
}

func (gateway *MysqlGateway) Update(item Item) (Item, error) {
	//
}

func (gateway *MysqlGateway) GetById(id uint) (Item, error) {
	var item Item

	rows, err := sq.Select("id, description, is_finished, created_at, finished_at").
		From(itemTable).
		Where(sq.Eq{"id": id}).
		RunWith(gateway.db).Query()

	if err != nil {
		return item, err
	}

	for rows.Next() {
		err = rows.Scan(&item.Id, &item.Description, &item.IsFinished, &item.CreatedAt, &item.FinishedAt)

		if err != nil {
			continue
		}

		return item, nil
	}

	rows.Close()

	return item, &NotFoundError{Message: "Item with Id " + strconv.Itoa(int(id)) + " cannot be found"}
}

func (gateway *MysqlGateway) GetAllFinished() ([]Item, error) {
	return gateway.getAll(true)
}

func (gateway *MysqlGateway) GetAllUnfinished() ([]Item, error) {
	return gateway.getAll(false)
}

func (gateway *MysqlGateway) getAll(isFinished bool) ([]Item, error) {
	var items []Item

	rows, err := sq.Select("id, description, is_finished, created_at, finished_at").
		From(itemTable).
		Where(sq.Eq{"is_finished": isFinished}).
		RunWith(gateway.db).Query()

	if err != nil {
		return items, err
	}

	var item Item
	for rows.Next() {
		err = rows.Scan(&item.Id, &item.Description, &item.IsFinished, &item.CreatedAt, &item.FinishedAt)

		if err != nil {
			continue
		}

		items = append(items, item)
	}

	rows.Close()

	return items, nil
}
