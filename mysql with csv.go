package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var myquery string

	csvFile, _ := os.Open("people.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	//I have created database named "mydb" with table "gotable"
	db, err := sql.Open("mysql", "talha:Talha1996@gmail.com@tcp(localhost)/mydb")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	for {
		line, error := reader.Read()
		if error == io.EOF {
			fmt.Println("\nAll lines read from csv")
			break
		} else if error != nil {
			log.Fatal(error)
		}
		
		myquery = "INSERT INTO gotable VALUES('" + line[0] + "','" + line[1] + "'," + line[2] + ",'" + line[3] + "')" + ";"
		fmt.Println("\n",line[0]," ",line[1],"\t",line[2],"\t",line[3])
		insert, err := db.Query(myquery)

		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()
	}

}
