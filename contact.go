//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede UG (haftungsbeschränkt) <info@jj-ideenschmiede.de>
//
// This file is part of sevdesk.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package sevdesk

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// To decode the data
type ContactReturn struct {
	Objects []ContactObject `json:"objects"`
}

type ContactObject struct {
	ID                        string      `json:"id"`
	ObjectName                string      `json:"objectName"`
	AdditionalInformation     interface{} `json:"additionalInformation"`
	Create                    string      `json:"create"`
	Update                    string      `json:"update"`
	Name                      string      `json:"name"`
	Status                    string      `json:"status"`
	CustomerNumber            string      `json:"customerNumber"`
	Surname                   string      `json:"surname"`
	Familyname                string      `json:"familyname"`
	Title                     string      `json:"title"`
	Category                  Category    `json:"category"`
	Description               string      `json:"description"`
	AcademicTitle             string      `json:"academicTitle"`
	Gender                    string      `json:"gender"`
	SevClient                 SevClient   `json:"sevClient"`
	Name2                     string      `json:"name2"`
	Birthday                  string      `json:"birthday"`
	VatNumber                 string      `json:"vatNumber"`
	BankAccount               interface{} `json:"bankAccount"`
	BankNumber                interface{} `json:"bankNumber"`
	DefaultCashbackTime       interface{} `json:"defaultCashbackTime"`
	DefaultCashbackPercent    interface{} `json:"defaultCashbackPercent"`
	DefaultTimeToPay          interface{} `json:"defaultTimeToPay"`
	TaxNumber                 interface{} `json:"taxNumber"`
	TaxOffice                 interface{} `json:"taxOffice"`
	ExemptVat                 string      `json:"exemptVat"`
	TaxType                   string      `json:"taxType"`
	DefaultDiscountAmount     interface{} `json:"defaultDiscountAmount"`
	DefaultDiscountPercentage interface{} `json:"defaultDiscountPercentage"`
	BuyerReference            interface{} `json:"buyerReference"`
	GovernmentAgency          string      `json:"governmentAgency"`
}

type Category struct {
	ID         string `json:"id"`
	ObjectName string `json:"objectName"`
}

type SevClient struct {
	ID         string `json:"id"`
	ObjectName string `json:"objectName"`
}

// To structure the data for new contact
type Contact struct {
	Name        string
	Name2       string
	Surname     string
	Familyname  string
	VatNumber   string
	TaxNumber   string
	BankAccount string
	BankNumber  string
	Category    string
	Token       string
}

// To decode the new contact return data
type NewContactReturn struct {
	Objects NewContactObjectReturn `json:"objects"`
}

type NewContactObjectReturn struct {
	ID                        string      `json:"id"`
	ObjectName                string      `json:"objectName"`
	AdditionalInformation     interface{} `json:"additionalInformation"`
	Create                    string      `json:"create"`
	Update                    string      `json:"update"`
	Name                      string      `json:"name"`
	Status                    string      `json:"status"`
	CustomerNumber            string      `json:"customerNumber"`
	Surname                   string      `json:"surname"`
	FamilyName                string      `json:"familyname"`
	Title                     string      `json:"title"`
	Category                  Category    `json:"category"`
	Description               string      `json:"description"`
	AcademicTitle             string      `json:"academicTitle"`
	Gender                    string      `json:"gender"`
	SevClient                 SevClient2  `json:"sevClient"`
	Name2                     string      `json:"name2"`
	Birthday                  string      `json:"birthday"`
	VatNumber                 string      `json:"vatNumber"`
	BankAccount               interface{} `json:"bankAccount"`
	BankNumber                interface{} `json:"bankNumber"`
	DefaultCashbackTime       interface{} `json:"defaultCashbackTime"`
	DefaultCashbackPercent    interface{} `json:"defaultCashbackPercent"`
	DefaultTimeToPay          interface{} `json:"defaultTimeToPay"`
	TaxNumber                 interface{} `json:"taxNumber"`
	TaxOffice                 interface{} `json:"taxOffice"`
	ExemptVat                 string      `json:"exemptVat"`
	TaxType                   string      `json:"taxType"`
	DefaultDiscountAmount     interface{} `json:"defaultDiscountAmount"`
	DefaultDiscountPercentage interface{} `json:"defaultDiscountPercentage"`
	BuyerReference            interface{} `json:"buyerReference"`
	GovernmentAgency          string      `json:"governmentAgency"`
}

type SevClient2 struct {
	ID                        string           `json:"id"`
	ObjectName                string           `json:"objectName"`
	AdditionalInformation     interface{}      `json:"additionalInformation"`
	Create                    string           `json:"create"`
	Update                    string           `json:"update"`
	Name                      string           `json:"name"`
	TemplateMainColor         string           `json:"templateMainColor"`
	TemplateSubColor          string           `json:"templateSubColor"`
	Status                    string           `json:"status"`
	AddressStreet             string           `json:"addressStreet"`
	AddressCity               string           `json:"addressCity"`
	AddressZip                string           `json:"addressZip"`
	MuncipalityKey            interface{}      `json:"muncipalityKey"`
	ContactPhone              string           `json:"contactPhone"`
	ContactEmail              string           `json:"contactEmail"`
	PayPalEmail               string           `json:"paypalEmail"`
	vatNumber                 string           `json:"vatNumber"`
	CeoName                   string           `json:"ceoName"`
	ContactFax                interface{}      `json:"contactFax"`
	Domain                    string           `json:"domain"`
	TemplateHeadlineColor     string           `json:"templateHeadlineColor"`
	Website                   string           `json:"website"`
	Bank                      string           `json:"bank"`
	bankNumber                string           `json:"bankNumber"`
	BankAccountNumber         string           `json:"bankAccountNumber"`
	DistrictCourt             string           `json:"districtCourt"`
	SmallSettlement           string           `json:"smallSettlement"`
	HasSeenBasicSettingsModal string           `json:"hasSeenBasicSettingsModal"`
	Bank2                     string           `json:"bank2"`
	BankIban                  string           `json:"bankIban"`
	BankBic                   string           `json:"bankBic"`
	OrderPaymentTerms         interface{}      `json:"orderPaymentTerms"`
	oderDeliveryTerms         interface{}      `json:"oderDeliveryTerms"`
	InvoiceTimeToPay          interface{}      `json:"invoiceTimeToPay"`
	Type                      string           `json:"type"`
	ForeignID                 interface{}      `json:"foreignId"`
	DefaultBillingTime        interface{}      `json:"defaultBillingTime"`
	BillingAccountNumber      interface{}      `json:"billingAccountNumber"`
	BillingBankCode           interface{}      `json:"billingBankCode"`
	CompanyRegister           string           `json:"companyRegister"`
	NameAddition              string           `json:"nameAddition"`
	ShowNet                   string           `json:"showNet"`
	PrintShowQr               string           `json:"printShowQr"`
	PrintShowPayPal           string           `json:"printShowPayPal"`
	PrintContactPerson        string           `json:"printContactPerson"`
	TaxNumber                 string           `json:"taxNumber"`
	PrintDeliveryReturn       string           `json:"printDeliveryReturn"`
	PrintPartNumber           string           `json:"printPartNumber"`
	PrintPosDescription       string           `json:"printPosDescription"`
	ChartOfAccounts           string           `json:"chartOfAccounts"`
	AccountingChart           AccountingChart  `json:"accountingChart"`
	AccountingSystem          AccountingSystem `json:"accountingSystem"`
	ContractNotePaymentTerms  interface{}      `json:"contractNotePaymentTerms"`
	ContractNoteDeliveryTerms interface{}      `json:"contractNoteDeliveryTerms"`
	PackingListDeliveryTerms  interface{}      `json:"packingListDeliveryTerms"`
	InvitedBy                 interface{}      `json:"invitedBy"`
	PartnerID                 interface{}      `json:"partnerId"`
	InviterRewarded           string           `json:"inviterRewarded"`
	InvoiceReminderDuration   interface{}      `json:"invoiceReminderDuration"`
	InvoiceReminderCharge     interface{}      `json:"invoiceReminderCharge"`
	AutosetDeliveryDate       string           `json:"autosetDeliveryDate"`
	AddressCountry            AddressCountry   `json:"addressCountry"`
	PrintPdfDeliveryReturn    string           `json:"printPdfDeliveryReturn"`
	PrintPdfFooter            string           `json:"printPdfFooter"`
	DefaultCurrency           string           `json:"defaultCurrency"`
	PrintPageNumbers          string           `json:"printPageNumbers"`
	FormOfCompany             string           `json:"formOfCompany"`
	CompanySize               string           `json:"companySize"`
	ReferralProgram           RreferralProgram `json:"referralProgram"`
	OnBoardingStatus          interface{}      `json:"onBoardingStatus"`
	PactasId                  string           `json:"pactasId"`
	DocServer                 string           `json:"docServer"`
	CreateInvoiceReminder     string           `json:"createInvoiceReminder"`
	AccountantNumber          string           `json:"accountantNumber"`
	AccountantClientNumber    string           `json:"accountantClientNumber"`
	AccountingYearBegin       string           `json:"accountingYearBegin"`
	CancelationDate           interface{}      `json:"cancelationDate"`
	FigoUsername              string           `json:"figoUsername"`
	SubscriptionStartDate     string           `json:"subscriptionStartDate"`
	SubscriptionEndDate       interface{}      `json:"subscriptionEndDate"`
	SubscriptionCycle         string           `json:"subscriptionCycle"`
	PrintFoldLines            string           `json:"printFoldLines"`
	DiscoveredAt              string           `json:"discoveredAt"`
	FormerBookkeepingTool     interface{}      `json:"formerBookkeepingTool"`
	HasAccountant             string           `json:"hasAccountant"`
	PartnerBrand              interface{}      `json:"partnerBrand"`
	PartnerCustomerId         interface{}      `json:"partnerCustomerId"`
	PartnerProvisioningId     interface{}      `json:"partnerProvisioningId"`
	PartnerWhitelabeled       string           `json:"partnerWhitelabeled"`
	Tenant                    string           `json:"tenant"`
	State                     string           `json:"state"`
}

type AccountingChart struct {
	ID         string `json:"id"`
	ObjectName string `json:"objectName"`
}

type AccountingSystem struct {
	ID         string `json:"id"`
	ObjectName string `json:"objectName"`
}

type AddressCountry struct {
	ID         string `json:"id"`
	ObjectName string `json:"objectName"`
}

type RreferralProgram struct {
	ID         string `json:"id"`
	ObjectName string `json:"objectName"`
}

// To structure the address data
type Address struct {
	Street    string
	Zip       string
	City      string
	ContactID string
	Token     string
}

// To decode the address return
type NewAddressReturn struct {
	Objects NewAddressObjectReturn `json:"objects"`
}

type NewAddressObjectReturn struct {
	ID                    string        `json:"id"`
	ObjectName            string        `json:"objectName"`
	AdditionalInformation interface{}   `json:"additionalInformation"`
	Create                string        `json:"create"`
	Update                string        `json:"update"`
	Contact               ContactObject `json:"contact"`
	Street                string        `json:"street"`
	Zip                   string        `json:"zip"`
	City                  string        `json:"city"`
	Country               Country       `json:"country"`
	Category              Category      `json:"category"`
	name                  interface{}   `json:"name"`
	SevClient             SevClient     `json:"sevClient"`
}

type Country struct {
	ID         string `json:"id"`
	ObjectName string `json:"objectName"`
}

// To structure the communication data
// The key represents what type it is (1: Private, 2: Work, 3. Fax, 4. Mobil, 5. empty, 6. Autobox, 7. Newsletter, 8. Invoice address)
type Communication struct {
	Key       string
	Value     string
	ContactID string
	Token     string
}

// To decode the phone data
type NewCommunicationReturn struct {
	Objects NewCommunicationObjectsReturn `json:"objects"`
}

type NewCommunicationObjectsReturn struct {
	ID                    string      `json:"id"`
	ObjectName            string      `json:"objectName"`
	AdditionalInformation interface{} `json:"additionalInformation"`
	Create                string      `json:"create"`
	Update                string      `json:"update"`
	Contact               Contact     `json:"contact"`
	Type                  string      `json:"type"`
	Value                 string      `json:"value"`
	Key                   Key         `json:"key"`
	Main                  string      `json:"main"`
	SevClient             SevClient   `json:"sevClient"`
}

type Key struct {
	ID         string `json:"id"`
	ObjectName string `json:"objectName"`
}

// Check contacts
func Contacts(token string) (ContactReturn, error) {

	// Define client
	client := &http.Client{}

	// New hhtp request
	request, err := http.NewRequest("GET", "https://my.sevdesk.de/api/v1/Contact", nil)
	if err != nil {
		return ContactReturn{}, err
	}

	// Set token to request header
	request.Header.Set("Authorization", token)

	// Response to sevDesk
	response, err := client.Do(request)
	if err != nil {
		return ContactReturn{}, err
	}

	// Close response
	defer response.Body.Close()

	// Decode data
	var decode ContactReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ContactReturn{}, err
	}

	// Return data
	return decode, nil

}

// Create new contact
func NewContact(config Contact) (NewContactReturn, error) {

	// Define new client
	client := &http.Client{}

	// Define body data
	body := url.Values{}

	body.Set("name", config.Name)
	body.Set("name2", config.Name2)
	body.Set("surname", config.Surname)
	body.Set("familyname", config.Familyname)
	body.Set("vatNumber", config.VatNumber)
	body.Set("taxNumber", config.TaxNumber)
	body.Set("bankAccount", config.BankAccount)
	body.Set("bankNumber", config.BankNumber)
	body.Set("category[id]", config.Category)
	body.Set("category[objectName]", "Category")

	// New http request
	request, err := http.NewRequest("POST", "https://my.sevdesk.de/api/v1/Contact", strings.NewReader(body.Encode()))
	if err != nil {
		return NewContactReturn{}, err
	}

	// Set header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", config.Token)

	// Response to sevDesk
	response, err := client.Do(request)
	if err != nil {
		return NewContactReturn{}, err
	}

	// Close response
	defer response.Body.Close()

	// Decode response body
	var decode NewContactReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return NewContactReturn{}, err
	}

	// Return the data
	return decode, nil

}

// Create new address
func NewAddress(config Address) (NewAddressReturn, error) {

	// Define new client
	client := &http.Client{}

	// Define body data
	body := url.Values{}

	body.Set("street", config.Street)
	body.Set("zip", config.Zip)
	body.Set("city", config.City)

	// New http request
	request, err := http.NewRequest("POST", "https://my.sevdesk.de/api/v1/Contact/"+config.ContactID+"/addAddress", strings.NewReader(body.Encode()))
	if err != nil {
		return NewAddressReturn{}, err
	}

	// Set header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", config.Token)

	// Response to sevDesk
	response, err := client.Do(request)
	if err != nil {
		return NewAddressReturn{}, err
	}

	// Close response
	defer response.Body.Close()

	// Decode response body
	var decode NewAddressReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return NewAddressReturn{}, err
	}

	// Return the data
	return decode, nil

}

// Create new phone
func NewPhone(config Communication) (NewCommunicationReturn, error) {

	// Create new communication
	phone, err := newCommunication(config, "/addPhone")
	if err != nil {
		return NewCommunicationReturn{}, err
	}

	// Return the data
	return phone, err

}

// Create new mobile
func NewMobile(config Communication) (NewCommunicationReturn, error) {

	// Create new communication
	mobile, err := newCommunication(config, "/addMobile")
	if err != nil {
		return NewCommunicationReturn{}, err
	}

	// Return the data
	return mobile, err

}

// Create new phone
func NewEmail(config Communication) (NewCommunicationReturn, error) {

	// Create new communication
	email, err := newCommunication(config, "/addEmail")
	if err != nil {
		return NewCommunicationReturn{}, err
	}

	// Return the data
	return email, err

}

// Create new website
func NewWebsite(config Communication) (NewCommunicationReturn, error) {

	// Create new communication
	website, err := newCommunication(config, "/addWeb")
	if err != nil {
		return NewCommunicationReturn{}, err
	}

	// Return the data
	return website, err

}

// To create new communication
func newCommunication(config Communication, add string) (NewCommunicationReturn, error) {

	// Define new client
	client := &http.Client{}

	// Define body data
	body := url.Values{}

	body.Set("key", config.Key)
	body.Set("value", config.Value)

	// New http request
	request, err := http.NewRequest("POST", "https://my.sevdesk.de/api/v1/Contact/"+config.ContactID+add, strings.NewReader(body.Encode()))
	if err != nil {
		return NewCommunicationReturn{}, err
	}

	// Set header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", config.Token)

	// Response to sevDesk
	response, err := client.Do(request)
	if err != nil {
		return NewCommunicationReturn{}, err
	}

	// Close response
	defer response.Body.Close()

	// Decode response body
	var decode NewCommunicationReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return NewCommunicationReturn{}, err
	}

	// Return the data
	return decode, nil

}
