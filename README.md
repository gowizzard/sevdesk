## sevDesk GO Lang
With this small library we want to simplify the handling of the sevDesk Api. Currently you can only create invoices and add invoices via ID single positions.

## Install

```console
go get github.com/jojojojonas/sevdesk
```

## How to use?
You can currently only create invoices and items in the invoices. For this purpose, sevDesk provides some parameters that are currently not well documented.

## Get invoices
It is now possible to read out the invoices. This is done in a very simple way. You use the following function with your user token and get back an object with slices.

```go
invoices, err := sevdesk.Invoices("token")
if err != nil {
fmt.Println("Error: ", err)
}

fmt.Println(invoices)
```

## Create invoice
Here you will find an example how to create a new invoice in sevDesk.

```go
invoice, err := sevdesk.NewInvoice(sevdesk.Invoice{"contactID", "address", "invoiceDate", "status", "invoiceType", "contactPerson", "subject", "headText", "footText", "token"})
if err != nil {
    fmt.Println("Error: ", err)
}
	
// Return the invoice id
fmt.Println(invoice.Objects.ID)
```

## Create new position in an invoice
Here is an example how to create a new position in an existing invoice.You need the ID of the invoice and of course the contact ID. You also need the parameter that distinguishes between pcs, hours or %.

For this I have worked it all out as follows. The funny thing is that the ID does not start with 0 but with 1:

**1 = Stk, 2 = m², 3 = m, 4 = kg, 5 = t, 6 = lfm, 7 = pauschal, 8 = m³, 9 = Std, 10 = km, 11 = %, 12 = Tag(e), 13 = l**

In sevDesk the price is transferred in gross, therefore we have added a calculation of the gross value to the function. So you set the net value + the VAT in the function.

```go
position, err := sevdesk.NewPosition(sevdesk.Position{"45", "1", "16", "Backups", "Backups of all Websites", "9", "invoiceID", "token"})
if err != nil {
    fmt.Println("Error: ", err)
}

// Return the invoice id
fmt.Println(position.Objects.ID)
```

## Send invoice via email
So that you can send your invoices directly there is now a new function. You can adjust several parameters like email, CC, BCC, subject and a text.

When using this function, an email is sent directly to the specified email address and the invoice is attached as a PDF.

```go
// Send email
email, err := sevdesk.SendInvoiceEmail(sevdesk.InvoiceEmail{invoice.Objects.ID, "email", "subject", "text", "cc", "bcc", "token"})
if err != nil {
    fmt.Println(err)
}
```

## Get contacts

If you want to read out all customers, then it goes as follows:

```go
contacts, err := sevdesk.Contacts("token")
if err != nil {
	fmt.Println("Error: ", err)
}
```

## Set new contact

If you want to set a new Contact, then this goes as follows. Some attributes are needed for this. The category[id] must be set. The best is a 3. Otherwise the other fields are free. The token must be specified of course.

```go
contact, err := sevdesk.NewContact(sevdesk.Contact{"Name", "Name2", "Surname", "Familyname", "Vat number", "Tax number", "Bank account", "Bank number", "CategoryID", "token"})
if err != nil {
	fmt.Println("Error: ", err)
}
```

## Add information

For each newly created contact, additional information about the contact can be added. Like an email, a phone number or a website. The CategorieID is very important. You can find it here.

**The key represents what type it is (1: Private, 2: Work, 3. Fax, 4. Mobil, 5. empty, 6. Autobox, 7. Newsletter, 8. Invoice address)**

### Add address

If you want to add an address, this is how to do it:

```go
address, err := sevdesk.NewAddress(sevdesk.Address{"Street", "Zip", "City", "ContactID", "Token"})
if err != nil {
	fmt.Println("Error: ", err)
}
```
 
### Add email

If you want to add an email, then this goes as follows:

```go
email, err := sevdesk.NewEmail(sevdesk.Communication{"Key", "Value", "ContactID", "Token"})
if err != nil {
    fmt.Println("Error: ", err)
}
```

### Add phone

If you want to add a phone number, this is how to do it:

```go
phone, err := sevdesk.NewPhone(sevdesk.Communication{"Key", "Value", "ContactID", "Token"})
if err != nil {
    fmt.Println("Error: ", err)
}
```

### Add mobile

If you want to add a mobile number, this is how to do it:

```go
mobile, err := sevdesk.NewMobile(sevdesk.Communication{"Key", "Value", "ContactID", "Token"})
if err != nil {
    fmt.Println("Error: ", err)
}
```

### Add website

If you want to add a website, this is how to do it:

```go
website, err := sevdesk.NewWebsite(sevdesk.Communication{"Key", "Value", "ContactID", "Token"})
if err != nil {
    fmt.Println("Error: ", err)
}
```

## Help
If you have any questions or comments, please contact us by e-mail at [jonas.kwiedor@jj-ideenschmiede.de](mailto:jonas.kwiedor@jj-ideenschmiede.de)