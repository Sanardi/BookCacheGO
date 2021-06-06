package main

import (
	"fmt"
)

// creating the struct for book, in real life this would probably be a databse elsewhere
type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title: \t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{
		ID:            1,
		Title:         "The Hitchhiker's guide to the galaxy",
		Author:        "Douglas Adams",
		YearPublished: 1979,
	},

	{
		ID:            2,
		Title:         "The Hobbit",
		Author:        "J.R.R> Tolkien",
		YearPublished: 1937,
	},

	{
		ID:            3,
		Title:         "Designing Data Intensive Applications",
		Author:        "Martin Kleppmann",
		YearPublished: 2016,
	},
}
