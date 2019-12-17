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
			Product: "Apple",
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
func TestCalculeteOrderConsistOf3MangosEquals60(t *testing.T) {
	result,err:=CalculeteOrder(Order{
		OrderLine{
			Product: "Mango",
			Count:   3,
		},
	})
	if err!=nil {
		t.Fatal("error")
	}
	if result!=90 {
		t.Fatal("result is incorrect")
	}
}
