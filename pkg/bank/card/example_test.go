package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	card := types.Card{
		Active:  true,
		Balance: 20_000_00,
		Name:    "InfinityWarSpecial",
		PAN:     "xxx-xxx-xxx-512",
	}
	PaymntSourc := PaymentSources(card)
	fmt.Println(PaymntSourc)
	//Output: {InfinityWarSpecial xxx-xxx-xxx-512 2000000}
}
