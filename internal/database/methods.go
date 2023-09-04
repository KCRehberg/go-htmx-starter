package database

import (
	"database/sql"

	"gorm.io/gorm"
)

func Create[T any](model *T) *gorm.DB {
	return DB.Create(&model)
}

func GetOne[T any](modelSearch *T) *gorm.DB {
	return DB.Where(&modelSearch).Find(&modelSearch)
}

func getQueryData(query string, params map[string]interface{}) (*sql.Rows, error) {
	if len(params) > 0 {
		return DB.Raw(query, params).Rows()
	} else {
		return DB.Raw(query).Rows()
	}
}

func GetQuery(query string, params map[string]interface{}) []map[string]interface{} {
	rows, err := getQueryData(query, params)
	checkErr(err)

	var columns []string
	columns, err = rows.Columns()
	checkErr(err)
	colNum := len(columns)

	var results []map[string]interface{}
	for rows.Next() {
		// Create interface
		record := make([]interface{}, colNum)
		for i := range record {
			record[i] = &record[i]
		}

		// Load data into record interface
		err = rows.Scan(record...)
		checkErr(err)

		// Map columns to record data
		var row = map[string]interface{}{}
		for i := range record {
			row[columns[i]] = record[i]
		}

		// Append to results
		results = append(results, row)
	}

	return results
}

func GetQueryFirst(query string, params map[string]interface{}) map[string]interface{} {
	data := GetQuery(query, params)

	if len(data) > 0 {
		return data[0]
	}

	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
