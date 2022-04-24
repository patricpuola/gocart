package cartservice_test

import (
	"fmt"
	"patricpuola/gocart/cartservice"
	"regexp"
	"testing"
)

type CustomerIdTest struct {
	customerId int
}

var cartTestCases = []CustomerIdTest{
	{124532},
	{54321},
	{794508},
	{176},
	{778658473},
}

const uuidRegex = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"

func TestNewCart(t *testing.T) {
	for _, customer := range cartTestCases {
		cart, err := cartservice.New(customer.customerId)
		if err != nil {
			t.Error(fmt.Sprintf("Failed to create a ShoppingCart with customerId %d", customer.customerId))
		}
		if cart.CustomerId != customer.customerId {
			t.Error(fmt.Sprintf("New Shoppingcart's CustomerId (%d) doesn't match (%d) given in creation", cart.CustomerId, customer.customerId))
		}
		if len(cart.Contents) > 0 {
			t.Error(fmt.Sprintf("New Shoppingcart created with CustomerId (%d) is not empty", customer.customerId))
		}
		uuidMatch, _ := regexp.MatchString(uuidRegex, cart.Uuid)
		if !uuidMatch {
			t.Error(fmt.Sprintf("New Shoppingcart's uuid created with CustomerId (%d) is not a valid Uuid (%s)", customer.customerId, cart.Uuid))
		}
	}
}
