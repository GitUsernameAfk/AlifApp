package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Active:  true,
			Balance: 20_000_00,
			Name:    "InfinityWarSpecial",
			PAN:     "xxx-xxx-xxx-512",
		},
		{
			Active:  false,
			Balance: 10_000_00,
			Name:    "AvatarSpecial",
			PAN:     "xxx-xxx-xxx-869",
		},
		{
			Active:  true,
			Balance: 0_00,
			Name:    "VISA",
			PAN:     "xxx-xxx-xxx-978",
		},
	}
	payment:=PaymentSources(cards)
	fmt.Println(payment)
	//Output: {InfinityWarSpecial xxx-xxx-xxx-512 2000000}
}
