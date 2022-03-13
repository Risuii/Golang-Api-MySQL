package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "risuii:babehlo123@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler((bookService))

	// ========================
	// FindAll Data Di Database
	// ========================

	// books, err := bookRepository.FindAll()

	// for _, book := range books {
	// 	fmt.Println("Title :", book.Title)
	// }

	// ========================
	// FindByID Data Di Database
	// ========================

	// book, err := bookRepository.FindByID(2)

	// fmt.Println("Title :", book.Title)

	// ========================
	// Memasukan Data Di Database
	// ========================

	// bookRequest := book.BookRequest{
	// 	Title:       "Gundam",
	// 	Description: "Good comic",
	// 	Price:       "200000",
	// 	Rating:      5,
	// 	Discount:    0,
	// }

	// bookRepository.Create(book)
	// bookService.Create(bookRequest)

	// =========================
	// CREATE DATA
	// =========================
	//
	// book := book.Book{}
	// book.Title = "Atomic Habits"
	// book.Price = 120000
	// book.Discount = 15
	// book.Rating = 4
	// book.Description = "ini adalah buku yang sangat bagus dari eka kurniawan"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")
	// }

	// =========================
	// READ DATA
	// =========================

	// var books []book.Book
	// var book book.Book
	// err = db.Debug().First(&book).Error
	// err = db.Debug().First(&book, 2).Error
	// err = db.Debug().Last(&book).Error
	// err = db.Debug().Find(&books).Error
	// err = db.Debug().Where("title = ?", "Man Tiger").Find(&books).Error
	// err = db.Debug().Where("rating = ?", 5).Find(&books).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title:", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	// fmt.Println("Title:", book.Title)
	// fmt.Println("book object %v", book)

	// =========================
	// UPDATE DATA
	// =========================

	// var books []book.Book
	// var book book.Book
	// err = db.Debug().First(&book).Error
	// err = db.Debug().First(&book, 2).Error
	// err = db.Debug().Last(&book).Error
	// err = db.Debug().Find(&books).Error
	// err = db.Debug().Where("title = ?", "Man Tiger").Find(&books).Error
	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }

	// book.Title = "Man Tiger (Revised edition)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Updating book record")
	// 	fmt.Println("==========================")
	// }

	// =========================
	// DELETE DATA
	// =========================

	// var book book.Book
	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Deleting  book record")
	// 	fmt.Println("==========================")
	// }

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetAllBooks)
	v1.GET("/books/:id", bookHandler.GetSingleBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	// v1.GET("/", bookHandler.RootHandler)
	// v1.GET("/hello", bookHandler.HelloHandler)
	// v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	// v1.GET("/query", bookHandler.QueryHandler)

	router.Run(":3000")
}

// urutan atau alur code

//repository diakses di dalam service
//service diakses di dalam handler

// urutan atau alur code ketika data masuk

// main
// handler
// serivce
// repository
// db
// mysql
