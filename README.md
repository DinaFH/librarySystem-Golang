# librarySystem-Golang
 consists of 4 files :
 
 a) client.go file :
 
 -which communicate with the librarian through command line asking him to choose what operation he want to be done on either on readers or books .
 
 it contains functions that send requests to the server based on the operation to be done :
 
 1.addReader()
 
 2.removeReader()
 
 3.getReaders()
 
 4.searchReaderByID()
 
 5.searchReaderByName()
 
 6.addBook() // try to add publication date of the book as year.
 
 7.getBooks()
 
 8.searchBookByID()
 
 9.searchBookByTitle()
 
 10.sortBookByTitle()
 
 11.sortBookByDate()
 
 b)server.go file :contains handlers of these functions

 c)readersData.txt :contains data of readers
 
 d)booksData.txt :contains data of books 
 
