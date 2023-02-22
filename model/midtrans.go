package model

import (
	"fmt"

	"github.com/google/uuid"
)

type MidtransData struct {
	typePayment string
}

func NewMidtransData(typePayment string) *MidtransData {
	return &MidtransData{
		typePayment: typePayment,
	}
}

func (m *MidtransData) InitDummyData() map[string]any {
	transactionDetailsContent := map[string]any{}
	transactionDetailsContent["order_id"] = fmt.Sprintf("order %v", uuid.New())
	transactionDetailsContent["gross_amount"] = 50000

	listData := []map[string]any{}
	itemDetailsContent := map[string]any{}
	itemDetailsContent["id"] = fmt.Sprintf("barang %v", uuid.New())
	itemDetailsContent["price"] = 50000
	itemDetailsContent["quantity"] = 1
	itemDetailsContent["name"] = "Barang dummy 1"

	listData = append(listData, itemDetailsContent)

	customerDetails := map[string]any{}
	customerDetails["first_name"] = "Dummy Midtrans"
	customerDetails["last_name"] = "User"
	customerDetails["email"] = "dummy@gmail.com"
	customerDetails["phone"] = "081223323423"

	gopayContent := map[string]any{}
	shopeepayContent := map[string]any{}
	if m.typePayment == "gopay" {
		gopayContent["enable_callback"] = true
		gopayContent["callback_url"] = "https://midtrans.com"
	} else if m.typePayment == "shopeepay" {
		shopeepayContent["enable_callback"] = true
		shopeepayContent["callback_url"] = "https://midtrans.com"
	}

	payload := map[string]any{}
	payload["payment_type"] = m.typePayment
	payload["transaction_details"] = transactionDetailsContent
	payload["item_details"] = listData
	payload["customer_details"] = customerDetails
	if len(gopayContent) != 0 {
		payload["gopay"] = gopayContent
	} else if len(shopeepayContent) != 0 {
		payload["shopeepay"] = shopeepayContent
	}
	return payload
}
