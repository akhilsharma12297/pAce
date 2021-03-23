package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func getConn() *sql.DB {

	var dbHost = "DB_HOST"
	var dbName = "DB_DATABASE"
	var dbPort = "DB_PORT"
	var dbUser = "DB_USER"
	var dbPassword = "DB_PASSWORD"
	var url = dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Print("Error opening SQL connection : ", err, " URL : ", url)
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		return nil
	}

	return db
}

func putTransaction(data Transaction) {

	db := getConn()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO POCKET_ACES_TRANSACTION (ID, AMOUNT, TYPE_VALUE, PARENT_ID) VALUES (?,?,?,?)")
	fmt.Println("Statement ", stmt)
	if stmt == nil {
		fmt.Println("ERROR : ", err)
	}

	_, err = stmt.Exec(data.transactionIDs, data.Amount, data.Type, data.ParentID)
	if err != nil {
		fmt.Println("Error creating row..")
	}

}

func getTransaction(transactionId string) (Transaction, bool) {

	query := `select ID, AMOUNT, TYPE_VALUE, PARENT_ID FROM POCKET_ACES_TRANSACTION WHERE ID = ?`

	db := getConn()
	defer db.Close()

	var resultData Transaction

	fmt.Println("Executing query : ", query)
	result, err := db.Query(query, transactionId)
	if err != nil {
		fmt.Println("Error creating row..")
		return resultData, false
	}
	if result == nil {
		fmt.Print("No rows read from DB")
		return resultData, false
	}

	for result.Next() {
		var pID sql.NullFloat64

		result.Scan(&resultData.transactionIDs, &resultData.Amount, &resultData.Type, &pID)

		resultData.ParentID = float32(pID.Float64)

	}

	return resultData, true

}

func getTransactionType(typeValue string) ([]string, bool) {

	query := `select ID FROM POCKET_ACES_TRANSACTION WHERE TYPE_VALUE = '` + typeValue + `'`

	db := getConn()
	defer db.Close()

	var resultData []string

	fmt.Println("Executing query : ", query)
	result, err := db.Query(query)
	if err != nil {
		fmt.Println("Error creating row..")
		return resultData, false
	}
	if result == nil {
		fmt.Print("No rows read from DB")
		return resultData, false
	}

	for result.Next() {
		var data sql.NullString
		_ = result.Scan(&data)
		resultData = append(resultData, data.String)
	}

	if len(resultData) == 0 {
		return resultData, false
	}

	return resultData, true

}

func getTransactionSum(transactionId string) (float64, bool) {

	data, isPresent := getTransaction(transactionId)

	var sum float64

	if isPresent {

		query := `select AMOUNT FROM POCKET_ACES_TRANSACTION WHERE ID <= '` + fmt.Sprintf("%f", data.ParentID) + `'`

		db := getConn()
		defer db.Close()

		fmt.Println("Executing query : ", query)
		result, err := db.Query(query)
		if err != nil {
			fmt.Println("Error creating row..")
			return sum, false
		}
		if result == nil {
			fmt.Print("No rows read from DB")
			return sum, false
		}

		for result.Next() {
			var val sql.NullFloat64
			result.Scan(&val)
			sum += val.Float64
		}

		return sum + data.Amount, true

	} else {
		return sum, false
	}

}
