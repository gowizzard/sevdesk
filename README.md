## sevDesk GO Lang
With this small library we want to simplify the handling of the sevDesk Api. Currently you can only create invoices and add invoices via ID single positions.

## Install

```console
go get github.com/jojojojonas/sevdesk
```

## How to use?
You can currently only create invoices and items in the invoices. For this purpose, sevDesk provides some parameters that are currently not well documented.

### Create invoice
Here you will find an example how to create a new invoice in sevDesk.

```go
invoice, err := sevdesk.NewInvoice(sevdesk.Invoice{"", "clientID", "21.11.2020", "100", "RE", "contactID", "token"})
	if err != nil {
		fmt.Println("Error: ", err)
	}
	
// Return the invoice id
fmt.Println(invoice.Objects.ID)
```

### Create new position in an invoice
Here is an example how to create a new position in an existing invoice.You need the ID of the invoice and of course the contact ID. You also need the parameter that distinguishes between pcs, hours or %.

For this I have worked it all out as follows. The funny thing is that the ID does not start with 0 but with 1:

**1 = Stk, 2 = m², 3 = m, 4 = kg, 5 = t, 6 = lfm, 7 = pauschal, 8 = m³, 9 = Std, 10 = km, 11 = %, 12 = Tag(e), 13 = l**

```go
Position, err := NewPosition(Position{"45", "1", "16", "Backups", "9", "invoiceID", "token"})
	if err != nil {
		fmt.Println("Error: ", err)
	}

// Return the invoice id
fmt.Println(position.Objects.ID)
```

## Help
If you have any questions or comments, please contact us by e-mail at [info@hilfebeiderwebsite.de](mailto:info@hilfebeiderwebsite.de)