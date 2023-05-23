package sql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Interface struct {
	Connection string
}

func (c *conn) getRecord(selection string, table string, query string) *sqlx.Rows {
	c.dbMu.Lock()
	defer c.dbMu.Unlock()

	result, err := c.db.Queryx(fmt.Sprintf(
		"select %s from %s where %s;",
		selection,
		table,
		query,
	))

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	result.Next()

	return result
}

func (c *conn) getRecords(table string, query string) [][]string {
	c.dbMu.Lock()
	defer c.dbMu.Unlock()

	record, err := c.db.Query(fmt.Sprintf(
		"select * from %s where %s;",
		table,
		query,
	))

	if err != nil {
		panic(err.Error())
	}
	defer record.Close()

	return nil
}

func (c *conn) getStruct(s interface{}, table string) {

}

func (c *conn) getTable(table string) [][]string {
	c.dbMu.Lock()
	defer c.dbMu.Unlock()

	record, err := c.db.Query(fmt.Sprintf(
		"select * from %s;",
		table,
	))

	if err != nil {
		panic(err.Error())
	}
	defer record.Close()

	return nil
}

func (c *conn) insertRecord(query string) {
	c.dbMu.Lock()
	defer c.dbMu.Unlock()

	insert, err := c.db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func (c *conn) insertRecords(queries []string) {
	c.dbMu.Lock()
	defer c.dbMu.Unlock()

	for _, query := range queries {
		insert, err := c.db.Query(query)

		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
	}
}

func (c *conn) insertStruct(s interface{}, table string) {

}

func (c *conn) updateRecord(table string, updates string, query string) {
	c.dbMu.Lock()
	defer c.dbMu.Unlock()

	update, err := c.db.Query(fmt.Sprintf(
		`update %s
		set %s
		where %s;`,
		table,
		updates,
		query,
	))

	if err != nil {
		panic(err.Error())
	}
	defer update.Close()
}

func (c *conn) updateRecords(table string, updates []string, queries []string) {
	c.dbMu.Lock()
	defer c.dbMu.Unlock()

	for index, element := range updates {
		update, err := c.db.Query(fmt.Sprintf(
			`update %s
			set %s
			where %s;`,
			table,
			element,
			queries[index],
		))

		if err != nil {
			panic(err.Error())
		}
		defer update.Close()
	}
}

func (c *conn) updateStruct(s interface{}, table string) {

}

func (c *conn) deleteRecord(table string, query string) {
	c.dbMu.Lock()
	defer c.dbMu.Unlock()

	delete, err := c.db.Query(fmt.Sprintf(
		"delete from %s where %s;",
		table,
		query,
	))

	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
}

func (c *conn) deleteRecords(table string, query string, records []string) { // TODO fix that
	query += records[0]
	for i := 1; i < len(records); i++ {
		query += ", " + records[i]
	}

	c.dbMu.Lock()
	defer c.dbMu.Unlock()

	delete, err := c.db.Query(fmt.Sprintf(
		"delete from %s where %s;",
		table,
		query,
	))

	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
}

func (c *conn) deleteStruct(s interface{}, table string) {
	
}

func (c *conn) deleteTable(table string) {
	c.dbMu.Lock()
	defer c.dbMu.Unlock()

	delete, err := c.db.Query(fmt.Sprintf(
		"drop table %s;",
		table,
	))

	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
}

func (c *conn) deleteTables(tables []string) {
	c.dbMu.Lock()
	defer c.dbMu.Unlock()

	for _, element := range tables {
		delete, err := c.db.Query(fmt.Sprintf(
			"drop table %s;",
			element,
		))

		if err != nil {
			panic(err.Error())
		}
		defer delete.Close()
	}
}
