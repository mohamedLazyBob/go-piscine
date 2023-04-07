package piscine

type food struct {
	orders   string
	preptime int
}

func SetFood(ptr *food) {
	if ptr.orders == "burger" {
		ptr.preptime = 15
	} else if ptr.orders == "chips" {
		ptr.preptime = 10
	} else if ptr.orders == "nuggets" {
		ptr.preptime = 12
	} else {
		ptr.preptime = 404
	}
}

func FoodDeliveryTime(order string) int {
	food := &food{}
	food.orders = order
	SetFood(food)
	return food.preptime
}
