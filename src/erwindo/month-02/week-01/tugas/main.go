package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type countDataCashier struct {
	count int
}

type session struct {
	name string
	id   int
}

type purchaseOrder struct {
	goodsName string
	quantity  int
	price     int
}

type lastTransaction struct {
	id         int
	grandTotal int
}

type transactionNumber struct {
	id   int
	date string
}

type monthlyReport struct {
	month     int
	total     int
	ammount   int
	purchased int
}

type transaction struct {
	id         int
	name       string
	date       string
	grandTotal int
}

type goodsPurchased struct {
	name  string
	qty   int
	price int
	total int
}

var po []purchaseOrder

var striped int = 27
var cashier = session{}

func main() {
	loginStatus := login()

	if loginStatus {
		clear()
		menuHome()
	} else {
		main()
	}
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:erwindo123@tcp(127.0.0.1:3306)/cashier")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func input() string {
	var kal string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	kal = scanner.Text()

	return kal
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func login() bool {

	fmt.Println(strings.Repeat("-", striped))
	fmt.Println("Cashier please login")
	fmt.Println(strings.Repeat("-", striped))

	fmt.Printf("Username: ")
	username := input()
	fmt.Printf("Password: ")
	password := input()

	logged, err := cekLogin(username, password)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if logged {
		return true
	}

	clear()
	fmt.Println("Error: incorrect username or password")
	return false

}

func cekLogin(username string, password string) (bool, error) {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	defer db.Close()

	var result = countDataCashier{}

	err = db.
		QueryRow("select count(*) as 'count' from cashiers where username = (?) and password = (?)", username, password).
		Scan(&result.count)

	createSession(username)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	if result.count > 0 {
		return true, nil
	}

	return false, nil
}

func createSession(username string) {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var result = session{}

	err = db.
		QueryRow("select id, name from cashiers where username = (?)", username).
		Scan(&result.id, &result.name)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cashier.id = result.id
	cashier.name = result.name

}

func menuHome() {
	menuChoosed := 0

	fmt.Println("Welcome,", cashier.name)
	fmt.Println()
	fmt.Println(strings.Repeat("-", striped))
	fmt.Println("Simple Cashier Application")
	fmt.Println(strings.Repeat("-", striped))
	fmt.Println("1. Input Transaction")
	fmt.Println("2. view the Report")
	fmt.Println()
	fmt.Println("3. Exit application")
	fmt.Println(strings.Repeat("-", striped))
	fmt.Println()

	for {
		fmt.Printf("Choose the Menu: ")
		menuChoosed, _ = strconv.Atoi(input())
		if menuChoosed > 0 && menuChoosed < 4 {
			break
		}
		fmt.Println("Error: you must insert 1-3")
	}

	switch menuChoosed {
	case 1:
		clear()
		menuInputTransaction()
		menuHome()
	case 2:
		clear()
		menuViewTransaction()
	case 3:
		clear()
		menuLogout()
	}

}

func menuLogout() {
	fmt.Println("Logout success")
	fmt.Println(strings.Repeat("-", striped))
	fmt.Println("Last admin login :", cashier.name)
	fmt.Println(strings.Repeat("-", striped))
	os.Exit(0)
}

func menuInputTransaction() {
	inputTransaction()
}

func inputTransaction() {
	fmt.Println("Welcome,", cashier.name)
	fmt.Println()
	fmt.Println(strings.Repeat("-", striped))
	fmt.Println("Create purchase order")
	fmt.Println(strings.Repeat("-", striped))
	fmt.Printf("Goods name : ")
	goodsName := input()
	fmt.Printf("Quantity   : ")
	quantity, _ := strconv.Atoi(input())
	fmt.Printf("Price (Rp) : ")
	price, _ := strconv.Atoi(input())

	po = append(po, purchaseOrder{goodsName: goodsName, quantity: quantity, price: price})

	fmt.Println(po)

	seeInsertedToStruct(goodsName, quantity, price)

	fmt.Printf("Do you want to continue? (y/n) : ")
	continueAction := input()

	if strings.ToLower(continueAction) == "y" {
		clear()
		inputTransaction()
	} else {
		clear()
		saveTransaction()
		menuHome()
	}

}

func saveTransaction() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	tx, err := db.Begin()
	handleError(err)

	currentTime := time.Now()

	res, err := tx.Exec("insert into transactions (cashier_id, transaction_date) VALUES(?, ?)", cashier.id, currentTime.Format("2006-01-02 15:04:05"))

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	handleError(err)

	for i := 0; i < len(po); i++ {
		res, err = tx.Exec("insert into purchases (name, qty, price, transaction_id) VALUES(?, ?, ?, ?)", po[i].goodsName, po[i].quantity, po[i].price, id)

		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	handleError(tx.Commit())

	po = nil

	grandTotalLastInserted(id)
	fmt.Println()
	log.Println("Saved")
	fmt.Println()
	fmt.Printf("Redirecting to home page...")
	time.Sleep(10 * time.Second)
	clear()
}

func seeInsertedToStruct(goodsName string, quantity int, price int) {
	clear()
	fmt.Println("Welcome,", cashier.name)
	fmt.Println()
	fmt.Println(strings.Repeat("-", striped))
	fmt.Println("Your purchased order")
	fmt.Println(strings.Repeat("-", striped))
	fmt.Println("Goods name  : ", goodsName)
	fmt.Println("Quantity    : ", quantity)
	fmt.Println("Price (Rp)  : ", price)
	fmt.Println(strings.Repeat("-", striped))
	fmt.Println("Total Price : ", "Rp.", quantity*price)
	fmt.Println()
}

func grandTotalLastInserted(id int64) {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var result = lastTransaction{}

	err = db.
		QueryRow("select transaction_id as id, sum(price) as grandTotal from purchases where transaction_id = (?) group by transaction_id", id).
		Scan(&result.id, &result.grandTotal)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Welcome,", cashier.name)
	fmt.Println()
	fmt.Println(strings.Repeat("-", striped))
	fmt.Println("Trx Number       : ", result.id)
	fmt.Println("Grand Total (Rp) : ", result.grandTotal)
	fmt.Println(strings.Repeat("-", striped))
}

func menuViewTransaction() {
	menuChoosed := 0

	fmt.Println("Welcome,", cashier.name)
	fmt.Println()
	fmt.Println(strings.Repeat("-", striped))
	fmt.Println("View the Report")
	fmt.Println(strings.Repeat("-", striped))
	fmt.Println("1. View All Trx Number")
	fmt.Println("2. View Transaction Detail")
	fmt.Println("3. View Monthly Report")
	fmt.Println()
	fmt.Println("4. Back to Home page")
	fmt.Println("5. Exit Application")
	fmt.Println(strings.Repeat("-", striped))
	fmt.Println()

	for {
		fmt.Printf("Choose the Menu: ")
		menuChoosed, _ = strconv.Atoi(input())
		if menuChoosed > 0 && menuChoosed < 6 {
			break
		}
		fmt.Println("Error: you must insert 1-5")
	}

	switch menuChoosed {
	case 1:
		clear()
		viewAllTrxNumber()
	case 2:
		clear()
		viewDetailTransaction()
	case 3:
		clear()
		viewMonthlyReport()
	case 4:
		clear()
		menuHome()
	case 5:
		clear()
		menuLogout()
	}

}

func viewAllTrxNumber() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, _ := db.Query("select id, transaction_date as date from transactions")

	defer rows.Close()

	var result []transactionNumber

	for rows.Next() {
		var each = transactionNumber{}
		var err = rows.Scan(&each.id, &each.date)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	fmt.Println("All Transaction number")
	fmt.Println("+------------------------------+")
	fmt.Println("| Trx ID | Date                |")
	fmt.Println("+------------------------------+")

	if len(result) < 1 {
		fmt.Println("| No data to displayed")
	}

	for _, each := range result {
		fmt.Println("|", each.id, space(len(strconv.Itoa(each.id)), 5), "|", each.date, "|")
	}

	fmt.Println("+------------------------------+")

	for {
		fmt.Println()
		fmt.Printf("(Press any key) to view report...")
		menuChoosed := input()
		if len(menuChoosed) > -1 {
			break
		}
	}

	clear()
	menuViewTransaction()
}

func viewDetailTransaction() {
	fmt.Printf("Insert a transaction code: ")
	trxID := input()

	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var result = transaction{}

	err = db.
		QueryRow("select tr.id as id, cs.name as name, tr.transaction_date as date, sum(po.price) as grandTotal from transactions tr join cashiers cs join purchases po on tr.cashier_id = cs.id and tr.id = po.transaction_id where tr.id = (?) group by tr.id", trxID).
		Scan(&result.id, &result.name, &result.date, &result.grandTotal)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	purcasedOrder, _ := db.Query("select name, qty, price, (qty*price) as total from purchases where transaction_id = (?)", result.id)

	defer purcasedOrder.Close()

	var goodsRow []goodsPurchased

	for purcasedOrder.Next() {
		var each = goodsPurchased{}
		var err = purcasedOrder.Scan(&each.name, &each.qty, &each.price, &each.total)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		goodsRow = append(goodsRow, each)
	}

	clear()
	fmt.Println("Welcome,", cashier.name)
	fmt.Println()
	fmt.Printf("+%v+\n", strings.Repeat("-", striped+36))
	fmt.Printf("Detail Trx of ID Trx : #%v\n", result.id)
	fmt.Printf("+%v+\n", strings.Repeat("-", striped+36))
	fmt.Println("Transaction code :", result.id)
	fmt.Println("Cashier name     :", result.name)
	fmt.Println("Transaction date :", result.date)
	fmt.Println("Grand Total      : Rp.", result.grandTotal)

	fmt.Printf("+%v+\n", strings.Repeat("-", striped+36))
	fmt.Println("| Goods name            | Quantity  | Price       | Total       |")
	fmt.Printf("+%v+\n", strings.Repeat("-", striped+36))

	for _, each := range goodsRow {
		fmt.Println("|", each.name, space(len(each.name), 20), "|", each.qty, space(len(strconv.Itoa(each.qty)), 8), "|", each.price, space(len(strconv.Itoa(each.price)), 10), "|", each.total, space(len(strconv.Itoa(each.total)), 10), "|")
	}

	fmt.Printf("+%v+\n", strings.Repeat("-", striped+36))

	for {
		fmt.Println()
		fmt.Printf("(Press any key) to view report...")
		menuChoosed := input()
		if len(menuChoosed) > -1 {
			break
		}
	}

	clear()
	menuViewTransaction()
}

func space(lenght, space int) string {
	output := ""
	for i := 0; i < space-lenght; i++ {
		output += " "
	}
	return output
}

func viewMonthlyReport() {
	fmt.Printf("Insert a year: ")
	year, _ := strconv.Atoi(input())

	clear()
	viewReport(year)
}

func viewReport(year int) {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, _ := db.Query("select extract(month from tr.transaction_date) as month, count(*) as total, sum(po.price) as ammount, count(po.id) as purchased from transactions tr join purchases po on tr.id = po.transaction_id where extract(year from tr.transaction_date) = (?) group by month", year)

	defer rows.Close()

	var result []monthlyReport

	for rows.Next() {
		var each = monthlyReport{}
		var err = rows.Scan(&each.month, &each.total, &each.ammount, &each.purchased)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	fmt.Printf("Purchased order in %v\n", year)
	fmt.Printf("+%v+\n", strings.Repeat("-", striped+20))
	fmt.Printf("| Month | Trx Total | Ammount value | Purchased |\n")
	fmt.Printf("+%v+\n", strings.Repeat("-", striped+20))

	if len(result) < 1 {
		fmt.Println("| No data to displayed")
	}

	for _, each := range result {
		fmt.Println("|", each.month, space(len(strconv.Itoa(each.month)), 4), "|", each.total, space(len(strconv.Itoa(each.total)), 8), "|", each.ammount, space(len(strconv.Itoa(each.ammount)), 12), "|", each.purchased, space(len(strconv.Itoa(each.purchased)), 8), "|")
	}
	fmt.Printf("+%v+\n", strings.Repeat("-", striped+20))

	for {
		fmt.Println()
		fmt.Printf("(Press any key) to view report...")
		menuChoosed := input()
		if len(menuChoosed) > -1 {
			break
		}
	}

	clear()
	menuViewTransaction()
}
