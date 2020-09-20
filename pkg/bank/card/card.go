package payment

import (
	"bank/pkg/github/types"
)

//PaymntS uses for sources from cards
func PaymntS(cards types.Card) types.PaymntSource {
	var PaymntSourc types.PaymntSource
	if cards.Active && cards.Balance > 0 {
		PaymntSourc.Number = cards.PAN
		PaymntSourc.Balance = cards.Balance
		PaymntSourc.Type = cards.Name
	}
	return PaymntSourc
}
