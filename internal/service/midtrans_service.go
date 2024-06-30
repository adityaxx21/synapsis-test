package service

import (
	"fmt"
	"math/rand/v2"
	// "synapsis-backend-test/internal/domain"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

// func GenerateSnapReq(user domain.User, item []domain.TransactionItem, transaction domain.Transaction ) *snap.Request {
// 	var items []midtrans.ItemDetails
// 	var price  int
// 	price = 0
// 	for _, d := range item {
// 		item := midtrans.ItemDetails{
// 			ID:    fmt.Sprintf("item-%d", d.ID),
// 			Price: int64(d.Price),
// 			Qty:   int32(d.Total),
// 			Name:  d.Title,
// 		}
// 		price += d.Price * d.Total
// 		items = append(items, item)
// 	}

// 	custAddress := &midtrans.CustomerAddress{
// 		FName: user.FName,
// 		LName: user.LName,
// 		Phone: user.Phone,
// 		Address: user.Address,
// 		City: user.City,
// 		Postcode: user.Postcode,
// 		CountryCode: "IDN",
// 	}

// 	// Initiate Snap Request
// 	snapReq := &snap.Request{
// 		TransactionDetails: midtrans.TransactionDetails{
// 			OrderID: fmt.Sprintf("ORDER-TRANSACTION-%d",transaction.ID),
// 			GrossAmt: int64(transaction.GrossAmount),
// 		}, 
// 		CreditCard: &snap.CreditCardDetails{
// 			Secure: true,
// 		},
// 		CustomerDetail: &midtrans.CustomerDetails{
// 			FName: user.FName,
// 			LName: user.LName,
// 			// Email: user.Email,
// 			Phone: user.Phone,
// 			BillAddr: custAddress,
// 			ShipAddr: custAddress,
// 		},
// 		Items: &items,
// 	}

//  return snapReq
// }


func GenerateSnapReqq() *snap.Request {

	// Initiate Customer address
	custAddress := &midtrans.CustomerAddress{
		FName:       "John",
		LName:       "Doe",
		Phone:       "081234567890",
		Address:     "Baker Street 97th",
		City:        "Jakarta",
		Postcode:    "16000",
		CountryCode: "IDN",
	}

	// Initiate Snap Request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  fmt.Sprintf("MID-GO-ID-%d", rand.IntN(100) ),
			GrossAmt: 200000,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    "John",
			LName:    "Doe",
			Email:    "john@doe.com",
			Phone:    "081234567890",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "ITEM1",
				Price: 200000,
				Qty:   1,
				Name:  "Someitem",
			},
		},
	}
	return snapReq
}