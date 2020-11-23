package sevdesk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Position struct {
	PriceNet    string
	Quantity    string
	TaxRate     string
	Name        string
	Description string
	UnityID     string
	InvoiceID   string
	Token       string
}

type PositionReturn struct {
	Objects PositionObjects `json:"objects"`
}

type PositionObjects struct {
	ID                         string                 `json:"id"`
	ObjectName                 string                 `json:"objectName"`
	AdditionalInformation      string                 `json:"additionalInformation"`
	Create                     string                 `json:"create"`
	Update                     string                 `json:"update"`
	Invoice                    InvoicePositionObjects `json:"invoice"`
	Quantity                   string                 `json:"quantity"`
	Price                      string                 `json:"price"`
	Name                       string                 `json:"name"`
	Priority                   string                 `json:"priority"`
	Unity                      ObjectName             `json:"unity"`
	SevClient                  ObjectName             `json:"sevClient"`
	PositionNumber             string                 `json:"positionNumber"`
	Text                       string                 `json:"text"`
	Discount                   string                 `json:"discount"`
	TaxRate                    string                 `json:"taxRate"`
	Temporary                  string                 `json:"temporary"`
	SumNet                     string                 `json:"sumNet"`
	SumGross                   string                 `json:"sumGross"`
	SumDiscount                string                 `json:"sumDiscount"`
	SumTax                     string                 `json:"sumTax"`
	SumNetAccounting           string                 `json:"sumNetAccounting"`
	SumTaxAccounting           string                 `json:"sumTaxAccounting"`
	SumGrossAccounting         string                 `json:"sumGrossAccounting"`
	PriceNet                   string                 `json:"priceNet"`
	PriceGross                 string                 `json:"priceGross"`
	PriceTax                   string                 `json:"priceTax"`
	SumNetForeignCurrency      string                 `json:"sumNetForeignCurrency"`
	SumTaxForeignCurrency      string                 `json:"sumTaxForeignCurrency"`
	SumGrossForeignCurrency    string                 `json:"sumGrossForeignCurrency"`
	SumDiscountForeignCurrency string                 `json:"sumDiscountForeignCurrency"`
	CreateNextPart             string                 `json:"createNextPart"`
}

type InvoicePositionObjects struct {
	Id                                  string     `json:"id"`
	ObjectName                          string     `json:"objectName"`
	AdditionalInformation               string     `json:"additionalInformation"`
	InvoiceNumber                       string     `json:"invoiceNumber"`
	Contact                             ObjectName `json:"contact"`
	Create                              string     `json:"create"`
	Update                              string     `json:"update"`
	SevClients                          ObjectName `json:"sev_clients"`
	InvoiceDate                         string     `json:"invoiceDate"`
	Header                              string     `json:"header"`
	HeadText                            string     `json:"headText"`
	FootText                            string     `json:"footText"`
	TimeToPay                           string     `json:"timeToPay"`
	DiscountTime                        string     `json:"discountTime"`
	Discount                            string     `json:"discount"`
	AddressName                         string     `json:"addressName"`
	AddressStreet                       string     `json:"addressStreet"`
	AddressZip                          string     `json:"addressZip"`
	AddressCity                         string     `json:"addressCity"`
	PayDate                             string     `json:"payDate"`
	CreateUser                          ObjectName `json:"createUser"`
	DeliveryDate                        string     `json:"deliveryDate"`
	Status                              string     `json:"status"`
	SmallSettlement                     string     `json:"smallSettlement"`
	ContactPerson                       ObjectName `json:"contactPerson"`
	TaxRate                             string     `json:"taxRate"`
	TaxText                             string     `json:"taxText"`
	DunningLevel                        string     `json:"dunningLevel"`
	AddressParentName                   string     `json:"addressParentName"`
	TaxType                             string     `json:"taxType"`
	SendDate                            string     `json:"sendDate"`
	OriginLastInvoice                   string     `json:"originLastInvoice"`
	InvoiceType                         string     `json:"invoiceType"`
	AccountIntervall                    string     `json:"accountIntervall"`
	AccountLastInvoice                  string     `json:"accountLastInvoice"`
	AccountNextInvoice                  string     `json:"accountNextInvoice"`
	ReminderTotal                       string     `json:"reminderTotal"`
	ReminderDebit                       string     `json:"reminderDebit"`
	ReminderDeadline                    string     `json:"reminderDeadline"`
	ReminderCharge                      string     `json:"reminderCharge"`
	AddressParentName2                  string     `json:"addressParentName2"`
	AddressName2                        string     `json:"addressName2"`
	AddressGender                       string     `json:"addressGender"`
	AccountStartDate                    string     `json:"accountStartDate"`
	AccountEndDate                      string     `json:"accountEndDate"`
	Address                             string     `json:"address"`
	Currency                            string     `json:"currency"`
	SumNet                              int        `json:"sumNet"`
	SumTax                              int        `json:"sumTax"`
	SumGross                            int        `json:"sumGross"`
	SumDiscounts                        int        `json:"sumDiscounts"`
	SumNetForeignCurrency               int        `json:"sumNetForeignCurrency"`
	SumTaxForeignCurrency               int        `json:"sumTaxForeignCurrency"`
	SumGrossForeignCurrency             int        `json:"sumGrossForeignCurrency"`
	SumDiscountsForeignCurrency         int        `json:"sumDiscountsForeignCurrency"`
	SumNetAccounting                    int        `json:"sumNetAccounting"`
	SumTaxAccounting                    int        `json:"sumTaxAccounting"`
	SumGrossAccounting                  int        `json:"sumGrossAccounting"`
	PaidAmount                          string     `json:"paidAmount"`
	CustomerInternalNote                string     `json:"customerInternalNote"`
	ShowNet                             string     `json:"showNet"`
	Enshrined                           string     `json:"enshrined"`
	SendType                            string     `json:"sendType"`
	DeliveryDateUntil                   string     `json:"deliveryDateUntil"`
	SendPaymentReceivedNotificationDate string     `json:"sendPaymentReceivedNotificationDate"`
	SumDiscountNet                      int        `json:"sumDiscountNet"`
	SumDiscountGross                    int        `json:"sumDiscountGross"`
	SumDiscountNetForeignCurrency       int        `json:"sumDiscountNetForeignCurrency"`
	SumDiscountGrossForeignCurrency     int        `json:"sumDiscountGrossForeignCurrency"`
}

// To create new
func NewPosition(config Position) (PositionReturn, error) {

	// Convert string to float64
	priceNet, err := strconv.ParseFloat(config.PriceNet, 64)
	if err != nil {
		return PositionReturn{}, err
	}

	// Convert taxRate from string to float64
	taxRate, err := strconv.ParseFloat(config.TaxRate, 64)
	if err != nil {
		return PositionReturn{}, err
	}

	// Calc gross
	priceGross := priceNet + (priceNet * taxRate / 100)

	// Define client
	client := &http.Client{}

	// Define body data
	body := url.Values{}
	body.Set("price", fmt.Sprintf("%.2f", priceGross))
	body.Set("quantity", config.Quantity)
	body.Set("taxRate", config.TaxRate)
	body.Set("name", config.Name)
	body.Set("text", config.Description)
	body.Set("unity[id]", config.UnityID)
	body.Set("unity[objectName]", "Unity")
	body.Set("objectName", "InvoicePos")
	body.Set("mapAll", "true")
	body.Set("invoice[id]", config.InvoiceID)
	body.Set("invoice[objectName]", "Invoice")

	// Define request
	request, err := http.NewRequest("POST", "https://my.sevdesk.de/api/v1/InvoicePos", strings.NewReader(body.Encode()))
	if err != nil {
		return PositionReturn{}, err
	}

	// Set header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", config.Token)

	// Response to sevdesk
	response, err := client.Do(request)
	if err != nil {
		return PositionReturn{}, err
	}

	// Close connection
	defer response.Body.Close()

	// Decode data
	var decode PositionReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return PositionReturn{}, err
	}

	// Return data
	return decode, nil

}
