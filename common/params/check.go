package params

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/libs/pkg/errors"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	tokentypes "github.com/okex/okchain-go-sdk/module/token/types"
	"strings"
)

const (
	tokenDescLenLimit = 256
	countDefault      = 100
	perPageDefault    = 50
	perPageMax        = 200
)

func CheckProduct(fromInfo keys.Info, passWd, product string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	if len(product) == 0 {
		return errors.New("failed. empty product")
	}

	return nil
}

func CheckDexAssets(fromInfo keys.Info, passWd, baseAsset, quoteAsset string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	if len(baseAsset) == 0 {
		return errors.New("failed. empty base asset")
	}

	if len(quoteAsset) == 0 {
		return errors.New("failed. empty quote asset")
	}

	return nil
}

// CheckQueryTokenInfo gives a quick validity check for the input params
func CheckQueryTokenInfo(ownerAddr, symbol string) error {
	if len(ownerAddr) == 0 && len(symbol) == 0 {
		return errors.New("failed. empty input")
	}

	return nil
}

func CheckTokenIssue(fromInfo keys.Info, passWd, orgSymbol, wholeName, tokenDesc string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	if len(orgSymbol) == 0 {
		return errors.New("failed. empty original symbol")
	}

	tokenDescLen := len(tokenDesc)
	if tokenDescLen == 0 || tokenDescLen > tokenDescLenLimit {
		return errors.New("failed. invalid token description")
	}

	if len(wholeName) == 0 {
		return errors.New("failed. empty whole name")
	}

	return nil
}

func CheckTransferUnitsParams(fromInfo keys.Info, passWd string, transfers []tokentypes.TransferUnit) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}
	transLen := len(transfers)
	if transLen == 0 {
		return errors.New("failed. no receiver input")
	}
	for i := 0; i < 0; i++ {
		if transfers[i].Coins.IsAllPositive() {
			continue
		} else {
			return errors.New("failed. only positive amount of coins is available")
		}
	}

	return nil
}

func CheckVoteParams(fromInfo keys.Info, passWd string, valAddrs []string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}
	if len(valAddrs) == 0 {
		return errors.New("failed. no validator address input")
	}

	// check duplicated
	filter := make(map[string]struct{}, len(valAddrs))
	for _, valAddr := range valAddrs {
		if _, ok := filter[valAddr]; ok {
			return fmt.Errorf("failed. validator address: %s is duplicated", valAddr)
		}
		filter[valAddr] = struct{}{}
	}

	return nil
}

func CheckKeyParams(fromInfo keys.Info, passWd string) error {
	if fromInfo == nil {
		return errors.New("failed. input invalid keys info")
	}
	if len(passWd) == 0 {
		return errors.New("failed. no password input")
	}

	return nil
}

func CheckSendParams(fromInfo keys.Info, passWd, toAddr string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}
	if len(toAddr) != 46 || !strings.HasPrefix(toAddr, "okchain") {
		return errors.New("failed. invalid receiver address")
	}

	return nil
}

func CheckNewOrderParams(fromInfo keys.Info, passWd string, products, sides, prices, quantities []string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	productsLen := len(products)
	if productsLen == 0 {
		return errors.New("failed. no product input")
	}

	if len(sides) != productsLen {
		return errors.New("failed. invalid param side counts")
	}

	if len(prices) != productsLen {
		return errors.New("failed. invalid param price counts")
	}

	if len(quantities) != productsLen {
		return errors.New("failed. invalid param quantity counts")
	}

	for _, side := range sides {
		if side != "BUY" && side != "SELL" {
			return errors.New(`failed. side must only be "BUY" or "SELL"`)
		}
	}

	return nil
}

func CheckCancelOrderParams(fromInfo keys.Info, passWd string, orderIds []string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	// check duplicated
	filter := make(map[string]struct{})
	for _, id := range orderIds {
		if _, ok := filter[id]; ok {
			return fmt.Errorf("failed. duplicated orderId: %s", id)
		}

		filter[id] = struct{}{}
	}

	return nil
}

// CheckQueryTickersParams gives a quick validity check for the input params of query tickers
func CheckQueryTickersParams(count []int) (countRet int, err error) {
	if len(count) > 1 {
		return countRet, errors.New("failed. invalid params input for tickers query")
	}

	if len(count) == 0 {
		countRet = countDefault
	} else {
		if count[0] < 0 {
			return countRet, errors.New(`failed. "count" is negative`)
		}
		countRet = count[0]
	}
	return
}

// CheckQueryRecentTxRecordParams gives a quick validity check for the input params of query recent tx record
func CheckQueryRecentTxRecordParams(product string, start, end, page, perPage int) (perPageRet int, err error) {
	if len(product) == 0 {
		return perPageRet, errors.New("failed. empty product")
	}

	return checkParamsPaging(start, end, page, perPage)
}

// CheckQueryOrdersParams gives a quick validity check for the input params of query orders
func CheckQueryOrdersParams(addrStr, product, side string, start, end, page, perPage int) (perPageRet int, err error) {
	if err = IsValidAccAddr(addrStr); err != nil {
		return
	}

	if len(product) == 0 {
		return perPageRet, errors.New("failed. empty product")
	}

	if !isValidSide(side) {
		return perPageRet, errors.New(`failed. "side" must only be "BUY" or "SELL"`)

	}

	return checkParamsPaging(start, end, page, perPage)
}

// CheckQueryTransactionsParams gives a quick validity check for the input params of query transactions
func CheckQueryTransactionsParams(addrStr string, typeCode, start, end, page, perPage int) (perPageRet int, err error) {
	if err = IsValidAccAddr(addrStr); err != nil {
		return
	}

	if typeCode < 0 {
		return perPageRet, errors.New("failed. type code isn't allowed to be negative")

	}

	return checkParamsPaging(start, end, page, perPage)
}

// IsValidAccAddr gives a quick validity check for an address string
func IsValidAccAddr(addrStr string) error {
	if len(addrStr) != 46 || !strings.HasPrefix(addrStr, "okchain") {
		return fmt.Errorf("failed. invalid account address: %s", addrStr)
	}
	return nil
}

func checkParamsPaging(start, end, page, perPage int) (perPageRet int, err error) {
	if start < 0 || end < 0 || page < 0 || perPage < 0 {
		return perPageRet, errors.New(`failed. "start","end","page","perPage" must be positive`)
	}

	if start > end {
		return perPageRet, errors.New(`failed. "start" isn't allowed to be larger than "end"`)
	}

	if perPage == 0 {
		perPageRet = perPageDefault
	} else if perPage > perPageMax {
		perPageRet = perPageMax
	} else {
		perPageRet = perPage
	}
	return

}

func isValidSide(side string) bool {
	if !(side == "BUY" || side == "SELL") {
		return false
	}
	return true
}
