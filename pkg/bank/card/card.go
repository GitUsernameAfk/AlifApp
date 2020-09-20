package card

import (
	"bank/pkg/bank/types"
)

//PaymentSources uses for sources from cards
func PaymentSources(cards types.Card) types.PaymentSource {
	var PaymntSourc types.PaymentSource
	if cards.Active && cards.Balance > 0 {
		PaymntSourc.Number = cards.PAN
		PaymntSourc.Balance = cards.Balance
		PaymntSourc.Type = cards.Name
	}
	return PaymntSourc
}
