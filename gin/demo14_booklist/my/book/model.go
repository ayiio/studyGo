package book

type Book struct {
	ID    int    `form:"id"`
	Title string `form:"title"`
	Price int    `form:"price"`
}

func ListBookFunc() (books []Book) {
	InitDB("test")
	books = ListBook()
	return
}

func (b *Book) NewBookFunc() {
	InitDB("test")
	SaveBook(b)
}

func (b *Book) DelBookFunc() {
	InitDB("test")
	DeleteBook(b.ID)
}
