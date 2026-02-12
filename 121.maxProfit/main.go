package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	// min_price := math.MaxInt
	// min_position, max_position := 0, 0
	// max_price := 0
	// max_profit := 0
	// for i, price := range prices {
	// 	if min_price > price {
	// 		min_price = price
	// 		min_position = i
	// 	}
	// 	if max_price < price {
	// 		max_price = price
	// 		max_position = i
	// 	}
	// 	if max_position < min_position {
	// 		max_price = min_price
	// 	} else {
	// 		if max_profit < max_price-min_price {
	// 			max_profit = max_price - min_price
	// 		}
	// 	}
	// }
	// return max(max_price-min_price, max_profit)

	buy := math.MaxInt
	sell := 0

	for _, price := range prices {
		buy = min(buy, price)
		sell = max(sell, price-buy)
	}
	return sell
}

func main() {
	test1 := []int{7, 1, 5, 3, 6, 4}
	ret := maxProfit(test1)
	fmt.Println(ret)

	test2 := []int{7, 6, 4, 3, 1}
	ret = maxProfit(test2)
	fmt.Println(ret)

	// ret = searchInsert(test2, 7)
	// fmt.Println(ret)

	ret = maxProfit([]int{2, 4, 1})
	fmt.Println(ret)
}
