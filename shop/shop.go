package shop

import "fmt"

type OrderLine struct {
	Product string
	Count int
}

type Order []OrderLine
var prices = map[string]int64 {
	"apple":20,
	"Banana":30,
}

// CalculeteOrder calculates grand total for order.
func CalculeteOrder(order Order) (int64, error) {
	var grandTotal int64
	for _,v:=range order {
		price,ok:=prices[v.Product]
		if !ok {
			return 0, fmt.Errorf("Price doesnt exsist")
		}
		grandTotal+=int64(int64(v.Count) * price)
	}
	return grandTotal, nil
}