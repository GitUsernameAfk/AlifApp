package payment

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExamplePaymntS() {
	card := types.Card{
		Active:  true,
		Balance: 20_000_00,
		Name:    "InfinityWarSpecial",
		PAN:     "xxx-xxx-xxx-512",
	}
	PaymntSourc := PaymntS(card)
	fmt.Println(PaymntSourc)
	//Output: {InfinityWarSpecial xxx-xxx-xxx-512 2000000}
}
