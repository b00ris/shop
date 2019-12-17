package shop

import "testing"

func TestCalculeteOrder(t *testing.T) {
	t.Skip("reason")
	result,_:=CalculeteOrder(Order{
		OrderLine{
			Product: "Apple",
			Count:   3,
		},
	})

	if result!=60 {
		t.Fatal("incorrect")
	}

}
func TestCalculeteOrderConsistOf3ApplesEquals60(t *testing.T) {
	result,err:=CalculeteOrder(Order{
		OrderLine{
			Product: "apple",
			Count:   3,
		},
	})
	if err!=nil {
		t.Fatal("error")
	}
	if result!=60 {
		t.FailNow()
	}
}
