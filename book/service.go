package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, book BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	// price dirubah karena tipe datanya json number
	price, _ := bookRequest.Price.Float64()
	rating, _ := bookRequest.Rating.Float64()
	discount, _ := bookRequest.Discount.Float64()
	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       int(price),
		Rating:      int(rating),
		Discount:    int(discount),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	// code dibawah untuk mencari data buku berdasarkan ID
	book, err := s.repository.FindByID(ID)
	price, _ := bookRequest.Price.Float64()
	rating, _ := bookRequest.Rating.Float64()
	discount, _ := bookRequest.Discount.Float64()

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Price = int(price)
	book.Rating = int(rating)
	book.Discount = int(discount)

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	// code dibawah untuk mencari data buku berdasarkan ID
	book, err := s.repository.FindByID(ID)

	newBook, err := s.repository.Delete(book)
	return newBook, err
}
