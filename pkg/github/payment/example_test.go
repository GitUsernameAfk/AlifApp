package payment
import(
	"fmt"
)
func ExamplePaymntS(){
	card := Card {
		Active:		true,
		Balance:	20_000_00,
		Name:		"InfinityWarSpecial",
		PAN:		"xxx-xxx-xxx-512",
	}
	PaymntSourc:=PaymntS(card)
	fmt.Println(PaymntSourc)
	//Output: {InfinityWarSpecial xxx-xxx-xxx-512 2000000}
}