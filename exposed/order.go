package exposed

import (
	"github.com/okex/okchain-go-sdk/crypto/keys"
	sdk "github.com/okex/okchain-go-sdk/types"
)

// Order shows the expected behavior for inner order client
type Order interface {
	sdk.Module
	OrderTx
	OrderQuery
}

// OrderTx shows the expected tx behavior for inner order client
type OrderTx interface {
	NewOrders(fromInfo keys.Info, passWd, products, sides, prices, quantities, memo string, accNum, seqNum uint64) (
		sdk.TxResponse, error)
}

// OrderQuery shows the expected query behavior for inner order client
type OrderQuery interface {
}