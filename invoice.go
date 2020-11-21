package sevdesk

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Invoice struct {
	InvoiceNumber string
	ContactID     string
	InvoiceDate   string
	Status        string
	InvoiceType   string
	ContactPerson string
	Token         string
}

type InvoiceReturn struct {
	Objects InvoiceObjects `json:"objects"`
}

type InvoiceObjects struct {
	ID                                  string     `json:"id"`
	ObjectName                          string     `json:"objectName"`
	AdditionalInformation               string     `json:"additionalInformation"`
	Contact                             ObjectName `json:"contact"`
	Create                              string     `json:"create"`
	Update                              string     `json:"update"`
	SevClient                           ObjectName `json:"sevClient"`
	InvoiceDate                         string     `json:"invoiceDate"`
	Header                              string     `json:"header"`
	HeadText                            string     `json:"headText"`
	FootText                            string     `json:"footText"`
	TimeToPlay                          string     `json:"timeToPlay"`
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
	SumNet                              string     `json:"sumNet"`
	SumTax                              string     `json:"sumTax"`
	SumGross                            string     `json:"sumGross"`
	SumDiscounts                        string     `json:"sumDiscounts"`
	SumNetForeignCurrency               string     `json:"sumNetForeignCurrency"`
	SumTaxForeignCurrency               string     `json:"sumTaxForeignCurrency"`
	SumGrossForeignCurrency             string     `json:"sumGrossForeignCurrency"`
	SumDiscountsForeignCurrency         string     `json:"sumDiscountsForeignCurrency"`
	SumNetAccounting                    string     `json:"sumNetAccounting"`
	SumTaxAccounting                    string     `json:"sumTaxAccounting"`
	SumGrossAccounting                  string     `json:"sumGrossAccounting"`
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

type ObjectName struct {
	Id         string `json:"id"`
	ObjectName string `json:"objectName"`
}

// Create a new
func NewInvoice(config Invoice) (InvoiceReturn, error) {

	// Define client
	client := &http.Client{}

	// Define the buddy
	body, err := json.Marshal(map[string]string{
		"invoiceNumber":             config.InvoiceNumber,
		"contact[id]":               config.ContactID,
		"contact[objectName]":       "Contact",
		"invoiceDate":               config.InvoiceDate,
		"header":                    "",
		"status":                    config.Status,
		"invoiceType":               config.InvoiceType,
		"currency":                  "EUR",
		"mapAll":                    "true",
		"objectName":                "Invoice",
		"discount":                  "false",
		"contactPerson[id]":         config.ContactPerson,
		"contactPerson[objectName]": "SevUser",
		"taxType":                   "default",
		"taxRate":                   "0",
		"taxText":                   "0",
		"showNet":                   "false",
	})
	if err != nil {
		return InvoiceReturn{}, err
	}

	// New http request
	request, err := http.NewRequest("POST", "https://my.sevdesk.de/api/v1/Invoice", bytes.NewBuffer(body))
	if err != nil {
		return InvoiceReturn{}, err
	}

	// Set header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", config.Token)

	// Response to sevDesk
	response, err := client.Do(request)

	// Close response
	defer response.Body.Close()

	// Decode data
	var decode InvoiceReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return InvoiceReturn{}, err
	}

	// Return data
	return decode, nil

}
