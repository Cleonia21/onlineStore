package main

type order struct {
	number   int
	products map[string]int
}

func newOrder() order {
	return order{
		number:   0,
		products: make(map[string]int),
	}
}

func (o *order) addProduct(name string, num int) {
	o.products[name] = num
}

type assemblyElement struct {
	productName string
	productId int
	productNum int
	orderId int
	additionalShelving []string
}

type collectionShelving  map[string] [] assemblyElement

func getOrder(ordersNums []int) collectionShelving {
	cs := make(collectionShelving)

}
