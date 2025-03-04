package items

type Product struct {
	Id    int
	Name  string
	Price float64
	Sku   string
}

type ItemList map[int]*Product

var (
	productIDs []int
	skuList    = make(map[int]*Product)
)

func New(id int) *Product {
	return &Product{}
}

func New() *Product {
	return &Product{}
}

func (il *ItemList) AddProductToItemList(p Product) {
	if p == 0 {
		p.Id = 1
		productIDs = append(productIDs, p.Id)
	}

}
