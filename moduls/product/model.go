package product

type Product struct {
	Id uint64
	Name string
	Price float64
}

var Products = []Product {
	Product { 1, "MacBook Pro 2020", 2400, },
}