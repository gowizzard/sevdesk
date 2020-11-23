package sevdesk

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// The data that the function uses
type Invoice struct {
	InvoiceNumber int
	ContactID     int
	InvoiceDate   string
	Status        int
	InvoiceType   string
	ContactPerson int
	Token         string
}

// For returning the data
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
	SumDiscountNet                      string     `json:"sumDiscountNet"`
	SumDiscountGross                    string     `json:"sumDiscountGross"`
	SumDiscountNetForeignCurrency       string     `json:"sumDiscountNetForeignCurrency"`
	SumDiscountGrossForeignCurrency     string     `json:"sumDiscountGrossForeignCurrency"`
}

// Uses for invoices and positions
type ObjectName struct {
	Id         string `json:"id"`
	ObjectName string `json:"objectName"`
}

// Create a new invoice
func NewInvoice(config Invoice) (InvoiceReturn, error) {

	// Define client
	client := &http.Client{}

	// Define body data
	body := url.Values{}
	body.Set("invoiceNumber", strconv.Itoa(config.InvoiceNumber))
	body.Set("contact[id]", strconv.Itoa(config.ContactID))
	body.Set("contact[objectName]", "Contact")
	body.Set("invoiceDate", config.InvoiceDate)
	body.Set("header", "")
	body.Set("status", strconv.Itoa(config.Status))
	body.Set("invoiceType", config.InvoiceType)
	body.Set("currency", "EUR")
	body.Set("mapAll", "true")
	body.Set("objectName", "Invoice")
	body.Set("discount", "false")
	body.Set("contactPerson[id]", strconv.Itoa(config.ContactPerson))
	body.Set("contactPerson[objectName]", "SevUser")
	body.Set("taxType", "default")
	body.Set("taxRate", "")
	body.Set("taxText", "0")
	body.Set("showNet", "false")

	// New http request
	request, err := http.NewRequest("POST", "https://my.sevdesk.de/api/v1/Invoice", strings.NewReader(body.Encode()))
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
