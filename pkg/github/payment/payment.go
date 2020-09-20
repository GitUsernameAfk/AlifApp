package payment

//Money is type for money (in lil units)
type Money int64

//Currency is type for currencies(TJS,USD,RUS etc.) 
type Currency string

//codes for currencies
const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUS Currency = "RUS"
)

//PAN is type for PAN of debet cards  
type PAN string

//PaymntSource is struct type fot sources from card 
type PaymntSource struct {
	Type 		string
	Number 		string
	Balance		Money
}


//Card is type for info of card
type Card struct{
	ID				int
	PAN				string
	Color			string
	Name			string
	Balance			Money		
	Currency		Currency
	Active			bool

}

//PaymntS uses for sources from cards
func PaymntS(cards Card) PaymntSource {
	var PaymntSourc PaymntSource
	if(cards.Active && cards.Balance>0) {
		PaymntSourc.Number = cards.PAN
		PaymntSourc.Balance=cards.Balance
		PaymntSourc.Type = cards.Name
	}
	return PaymntSourc
}
