package card

import (
	"bank/pkg/bank/types"
)

//PaymentSources uses for sources from cards
func PaymentSources(cards []types.Card) []types.PaymentSource {
	var PaymntSources []types.PaymentSource
	for _, card := range cards {

		if card.Active && card.Balance > 0 {
			PaymntSources = append(PaymntSources, 
				types.PaymentSource {
					Number:		card.PAN,
					Balance:	card.Balance,
					Type:		card.Name,
				} )
		}
	}
	return PaymntSources
}

