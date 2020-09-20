package types

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

//Card is type for info of card
type Card struct {
	ID       int
	PAN      PAN
	Color    string
	Name     string
	Balance  Money
	Currency Currency
	Active   bool
}

//Payment uses for payments of cards
type Payment struct {
	ID     int64
	Amount Money
}

//PaymentSource is struct type fot sources from card
type PaymentSource struct {
	Type    string
	Number  PAN
	Balance Money
}
