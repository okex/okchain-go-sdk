package utils

import (
	"encoding/json"
	ordertypes "github.com/okex/okexchain/x/order/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetOrderIDsFromResponse(t *testing.T) {
	mockOrderIDs := []string{"ID0000000000-1", "ID0000000000-2", "ID0000000000-3", "ID0000000000-4", "ID0000000000-5"}
	var orderResults, fakeOrderResults1, fakeOrderResults2 []ordertypes.OrderResult
	for i, orderID := range mockOrderIDs {
		mockOrderRes := buildMockOrderRes(orderID)
		if i < 3 {
			orderResults = append(orderResults, mockOrderRes)
		} else if i == 3 {
			fakeOrderResults1 = append(fakeOrderResults1, mockOrderRes)
		} else {
			fakeOrderResults2 = append(fakeOrderResults2, mockOrderRes)
		}
	}

	rawStrs := getRawStrSlice(orderResults, fakeOrderResults1, fakeOrderResults2)
	rawStrs = append(rawStrs, "string that failed to unmarshal JSON")
	require.Equal(t, 4, len(rawStrs))

	// build mock TxResponse
	mockTxResp := sdk.TxResponse{
		Logs: sdk.ABCIMessageLogs{
			{
				MsgIndex: 0,
				Log:      "default log",
				Events: sdk.StringEvents{
					{
						Type: "message",
						Attributes: []sdk.Attribute{
							{
								Key:   "not orders",
								Value: rawStrs[1],
							},
							{
								Key:   "orders",
								Value: rawStrs[3], // log error
							},
							{
								Key:   "orders",
								Value: rawStrs[0], // target
							},
						},
					},
					{
						Type: "not message",
						Attributes: []sdk.Attribute{
							{
								Key:   "not orders",
								Value: rawStrs[1],
							},
							{
								Key:   "orders",
								Value: rawStrs[2],
							},
						},
					},
				},
			},
		},
	}

	orderIDs, err := GetOrderIDsFromResponse(&mockTxResp)
	require.NoError(t, err)
	require.Equal(t, mockOrderIDs[:3], orderIDs)
}

func getRawStrSlice(orderResults ...[]ordertypes.OrderResult) (strs []string) {
	for _, res := range orderResults {
		bytes, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		strs = append(strs, string(bytes))
	}

	return
}

func buildMockOrderRes(orderID string) ordertypes.OrderResult {
	return ordertypes.OrderResult{
		Message: "default message",
		OrderID: orderID,
	}
}
