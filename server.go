package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

type Reader struct {
	ID                                 int
	Name, Gender, Birthday, Employment string
	Weight, Height                     float64
}
type Book struct {
	ID                                                         int
	Title, PublicationDate, Author, Genre, Publisher, Language string
}

var readerPath = "readersData.txt"
var bookPath = "booksData.txt"

//Reader operations handlers
// add a reader handler
func addReaderHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, _ := ioutil.ReadAll(request.Body)
		var reader Reader
		json.Unmarshal(body, &reader) //mapping from json to struct
		file, err1 := os.OpenFile(readerPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err1 != nil {
			panic(err1)
		}
		_, err := file.WriteString(string(body) + "\n")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		fmt.Fprintf(response, "Received!")
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////
// remove a reader handler
func removeReaderHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {

		keys, ok := request.URL.Query()["id"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'id' is missing")
			return
		}
		id, _ := strconv.Atoi(keys[0])

		file, err1 := os.Open(readerPath)
		if err1 != nil {
			panic(err1)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var readerArray []Reader
		for scanner.Scan() {
			var reader Reader
			json.Unmarshal([]byte(scanner.Text()), &reader)
			readerArray = append(readerArray, reader)
		}
		err2 := file.Close()
		if err2 != nil {
			panic(err2)
		}

		err3 := os.Remove(readerPath)
		if err3 != nil {
			panic(err3)
		}

		file, err4 := os.Create(readerPath)
		if err4 != nil {
			panic(err4)
		}
		file, err5 := os.OpenFile(readerPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err5 != nil {
			panic(err5)
		}
		for i, member := range readerArray {
			if member.ID == id {
				readerArray = append(readerArray[:i], readerArray[i+1:]...)
			}
		}
		for _, member := range readerArray {
			line, _ := json.Marshal(member)
			_, err6 := file.WriteString(string(line) + "\n")
			if err6 != nil {
				panic(err6)
			}
		}
		fmt.Println("Done")
		defer file.Close()

		fmt.Fprintf(response, "Reader has been removed!")

	}
}

////////////////////////////////////////////////
// get readers handler
func getReadersHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {

		file, err1 := os.Open(readerPath)
		if err1 != nil {
			panic(err1)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var readerArray []Reader
		for scanner.Scan() {
			//fmt.Fprintf(response, "Reader:", scanner.Text())
			var reader Reader
			json.Unmarshal([]byte(scanner.Text()), &reader)
			readerArray = append(readerArray, reader)
		}

		for _, member := range readerArray {

			fmt.Fprintf(response, "Reader:", fmt.Sprint(member))
		}
	}
}

////////////////////////////////////////////////////////////////////////
//search reader by id handler
func searchReaderIDHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {

		keys, ok := request.URL.Query()["id"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'id' is missing")
			return
		}
		id, _ := strconv.Atoi(keys[0])

		file, err1 := os.Open(readerPath)
		if err1 != nil {
			panic(err1)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var readerArray []Reader
		for scanner.Scan() {
			var reader Reader
			json.Unmarshal([]byte(scanner.Text()), &reader)
			readerArray = append(readerArray, reader)
		}

		for _, member := range readerArray {
			if member.ID == id {
				fmt.Fprintf(response, "Reader is found and this is his info:", fmt.Sprint(member))

			}

		}
	}
}

/////////////////////////////////////////////////////////////////////
//search reader by name handler
func searchReaderNameHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {

		keys, ok := request.URL.Query()["name"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'name' is missing")
			return
		}

		name := keys[0]

		file, err1 := os.Open(readerPath)
		if err1 != nil {
			panic(err1)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var readerArray []Reader
		for scanner.Scan() {
			var reader Reader
			json.Unmarshal([]byte(scanner.Text()), &reader)
			readerArray = append(readerArray, reader)
		}

		for _, member := range readerArray {
			if member.Name == name {
				fmt.Fprintf(response, "Reader is found and this is his info", fmt.Sprint(member))

			}
		}
	}
}

/////////////////////////////////////////////////////////////
//Book operations handlers:
// add book handler
func addBookHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, _ := ioutil.ReadAll(request.Body)
		var book Book
		json.Unmarshal(body, &book) //mapping from json to struct
		file, err1 := os.OpenFile(bookPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err1 != nil {
			panic(err1)
		}
		_, err := file.WriteString(string(body) + "\n")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		fmt.Fprintf(response, "Received!")
	}
}

//////////////////////////////////////////////////////////////////////////////////////
//get books handler

func getBooksHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {

		file, err1 := os.Open(bookPath)
		if err1 != nil {
			panic(err1)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var bookArray []Book
		for scanner.Scan() {
			//fmt.Fprintf(response, "Reader:", scanner.Text())
			var book Book
			json.Unmarshal([]byte(scanner.Text()), &book)
			bookArray = append(bookArray, book)
		}

		for _, member := range bookArray {

			fmt.Fprintf(response, "Book:", fmt.Sprint(member))
		}
	}
}

/////////////////////////////////////////////////////////
//search book by id handler
func searchBookIDHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {

		keys, ok := request.URL.Query()["id"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'id' is missing")
			return
		}
		id, _ := strconv.Atoi(keys[0])

		file, err1 := os.Open(bookPath)
		if err1 != nil {
			panic(err1)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var bookArray []Book
		for scanner.Scan() {
			var book Book
			json.Unmarshal([]byte(scanner.Text()), &book)
			bookArray = append(bookArray, book)
		}

		for _, member := range bookArray {
			if member.ID == id {
				fmt.Fprintf(response, "Book is found and this is its ifno :", fmt.Sprint(member))

			}
		}
	}
}

//////////////////////////////////////////////////////////////
//search book by title handler
func searchBookTitleHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {

		keys, ok := request.URL.Query()["title"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'title' is missing")
			return
		}

		title := keys[0]

		file, err1 := os.Open(bookPath)
		if err1 != nil {
			panic(err1)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var bookArray []Book
		for scanner.Scan() {
			var book Book
			json.Unmarshal([]byte(scanner.Text()), &book)
			bookArray = append(bookArray, book)
		}

		for _, member := range bookArray {
			if member.Title == title {
				fmt.Fprintf(response, "Book is found and this is its info", fmt.Sprint(member))
			}

		}
	}
}

////////////////////////////////////////////////////////////
// serach book by title  handler
func sortByTileHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {

		file, err1 := os.Open(bookPath)
		if err1 != nil {
			panic(err1)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var bookArray []Book
		for scanner.Scan() {
			//fmt.Fprintf(response, "Reader:", scanner.Text())
			var book Book
			json.Unmarshal([]byte(scanner.Text()), &book)
			bookArray = append(bookArray, book)
		}
		sort.SliceStable(bookArray, func(i, j int) bool {
			return bookArray[i].Title < bookArray[j].Title
		})

		for _, member := range bookArray {

			fmt.Fprintf(response, "Books sorted by title:", fmt.Sprint(member))
		}
	}
}

//////////////////////////////////////////////////////////////////////
// search book by publication date handler
func sortByDateHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {

		file, err1 := os.Open(bookPath)
		if err1 != nil {
			panic(err1)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var bookArray []Book
		for scanner.Scan() {
			//fmt.Fprintf(response, "Reader:", scanner.Text())
			var book Book
			json.Unmarshal([]byte(scanner.Text()), &book)
			bookArray = append(bookArray, book)
		}
		sort.SliceStable(bookArray, func(i, j int) bool {
			return bookArray[i].PublicationDate < bookArray[j].PublicationDate
		})

		for _, member := range bookArray {

			fmt.Fprintf(response, "Books sorted by date:", fmt.Sprint(member))
		}
	}
}

///////////////////////////////////////////////////
func main() {
	//Reader Handlers
	http.HandleFunc("/reader/add", addReaderHandler)
	http.HandleFunc("/reader/remove", removeReaderHandler)
	http.HandleFunc("/reader/get", getReadersHandler)
	http.HandleFunc("/reader/searchid", searchReaderIDHandler)
	http.HandleFunc("/reader/searchname", searchReaderNameHandler)
	// Book Handlers
	http.HandleFunc("/book/add", addBookHandler)
	http.HandleFunc("/book/get", getBooksHandler)
	http.HandleFunc("/book/searchid", searchBookIDHandler)
	http.HandleFunc("/book/searchtitle", searchBookTitleHandler)
	http.HandleFunc("/book/sorttitle", sortByTileHandler)
	http.HandleFunc("/book/sortdate", sortByDateHandler)

	err := http.ListenAndServe(":8050", nil)
	if err != nil {
		panic(err)
	}

}
