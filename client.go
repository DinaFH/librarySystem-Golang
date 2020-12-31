package main

import (
	
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

// Reader operations

func addReader() {
	var id int
	var name, gender, birthday, employment string
	var weight, height float64
	println("Please Enter Reader information")
	println("Enter id of reader ")
	fmt.Scanln(&id)
	println("Enter name of reader ")
	fmt.Scanln(&name)
	println("Enter gender of reader ")
	fmt.Scanln(&gender)
	println("Enter birthday of reader ")
	fmt.Scanln(&birthday)
	println("Enter job of reader ")
	fmt.Scanln(&employment)
	println("Enter weight of reader ")
	fmt.Scanln(&weight)
	println("Enter height of reader ")
	fmt.Scanln(&height)
	reader := Reader{
		ID:         id,
		Name:       name,
		Gender:     gender,
		Birthday:   birthday,
		Employment: employment,
		Weight:     weight,
		Height:     height,
	}
	r, _ := json.Marshal(reader)

	response, err := http.Post("http://localhost:8050/reader/add", "", bytes.NewBuffer([]byte(r)))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}

////////////////////////////////////////////////////////////////////
func removeReader() {
	var id int
	println("Enter id of reader ")
	fmt.Scanln(&id)

	response, err := http.Get("http://localhost:8050/reader/remove?id=" + strconv.Itoa(id))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}

/////////////////////////////////////////////////////////////////////////////////////////////////
func getReaders() {

	response, err := http.Get("http://localhost:8050/reader/get")

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}

///////////////////////////////////////////////////////////
func searchReaderByID() {
	var id int
	println("Enter id of reader ")
	fmt.Scanln(&id)

	response, err := http.Get("http://localhost:8050/reader/searchid?id=" + strconv.Itoa(id))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}

///////////////////////////////////////////////////////////////////////
func searchReaderByName() {
	var name string
	println("Enter name of reader ")
	fmt.Scanln(&name)

	response, err := http.Get("http://localhost:8050/reader/searchname?name=" + name)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}

///////////////////////////////////////////////////////////////
//Book operations:

func addBook() {
	var id int

	var title, publicationDate, author, genre, publisher, language string

	println("Please Enter Book information")
	println("Enter id the of book ")
	fmt.Scanln(&id)
	println("Enter title of the book ")
	fmt.Scanln(&title)
	println("Enter publication date of the book")
	fmt.Scanln(&publicationDate)
	println("Enter author of the book")
	fmt.Scanln(&author)
	println("Enter genre of the book ")
	fmt.Scanln(&genre)
	println("Enter publisher of the book")
	fmt.Scanln(&publisher)
	println("Enter language of the book")
	fmt.Scanln(&language)
	book := Book{
		ID:              id,
		Title:           title,
		PublicationDate: publicationDate,
		Author:          author,
		Genre:           genre,
		Publisher:       publisher,
		Language:        language,
	}
	b, _ := json.Marshal(book)

	response, err := http.Post("http://localhost:8050/book/add", "", bytes.NewBuffer([]byte(b)))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}

/////////////////////////////////////////////////////////////
func getBooks() {

	response, err := http.Get("http://localhost:8050/book/get")

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}

///////////////////////////////////////////////////////////////////
func searchBookByID() {
	var id int
	println("Enter id of the book")
	fmt.Scanln(&id)

	response, err := http.Get("http://localhost:8050/book/searchid?id=" + strconv.Itoa(id))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

///////////////////////////////////////////////////////////////////
func searchBookByTitle() {
	var title string
	println("Enter title of reader ")
	fmt.Scanln(&title)

	response, err := http.Get("http://localhost:8050/book/searchtitle?title=" + title)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}

//////////////////////////////////////////////////////////////
func sortBookByTitle() {

	response, err := http.Get("http://localhost:8050/book/sorttitle")

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}

/////////////////////////////////////////////////////////////
func sortBookByDate() {

	response, err := http.Get("http://localhost:8050/book/sortdate")

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}
func main() {
	var operation string
	var help string

	fmt.Println("Enter 'h' for displaying help")
	fmt.Scanln(&help)

	result := help == "h"
	fmt.Println("'ar' : add a reader")
	fmt.Println("'rr' : remove a reader")
	fmt.Println("'gr' : get all readers info ")
	fmt.Println("'ri' : search for a reader by id")
	fmt.Println("'rn' : search for a reader by name")
	fmt.Println("'ab' : add a book")
	fmt.Println("'gb' : get all books info ")
	fmt.Println("'bi' : search for a book by id")
	fmt.Println("'bt' : search for a book by title")
	fmt.Println("'st': sort a book by title")
	fmt.Println("'sd': sort a book by publication date")

	if result {
		for true {

			fmt.Scanln(&operation)

			if operation == "ar" {
				addReader()

			} else if operation == "rr" {
				removeReader()

			} else if operation == "gr" {
				getReaders()

			} else if operation == "ri" {
				searchReaderByID()

			} else if operation == "rn" {
				searchReaderByName()

			} else if operation == "ab" {
				addBook()

			} else if operation == "gb" {
				getBooks()

			} else if operation == "bi" {

				searchBookByID()

			} else if operation == "bt" {

				searchBookByTitle()
			} else if operation == "st" {

				sortBookByTitle()
			} else if operation == "sd" {

				sortBookByDate()
			} else {
				fmt.Println("Please insert a valid option")

			}
		}

	}

}
