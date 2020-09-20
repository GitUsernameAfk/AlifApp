package card

import (
	"bank/pkg/bank/types"
)

//PaymentSources uses for sources from cards
func PaymentSources(cards []types.Card) types.PaymentSource {
	var PaymntSourc types.PaymentSource
	for _, card := range cards {
	if card.Active && card.Balance > 0 {
		PaymntSourc.Number = card.PAN
		PaymntSourc.Balance = card.Balance
		PaymntSourc.Type = card.Name
	}
	}
	return PaymntSourc
}
