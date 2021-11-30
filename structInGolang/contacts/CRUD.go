package contacts

import (
	"fmt"
	"os"
)

type Contact struct {
	ID        uint8
	FirstName string
	LastName  string
	Phone     int64
	Email     string
	Position  string
}

var contacts []Contact

func (*Contact) GetAll() {

	for _, contact := range contacts {

		fmt.Printf("\nId:%d\nFirst Name:%s\nLast Name:%s\nPhone:%d\nEmail:%s\nPosition:%s\n\n",
			contact.ID,
			contact.FirstName,
			contact.LastName,
			contact.Phone,
			contact.Email,
			contact.Position,
		)
	}
}

func (*Contact) Get() {

	var id uint8

	fmt.Println("Enter contact ID:")
	fmt.Scan(&id)

	fmt.Printf("\nId:%d\nFirst Name:%s\nLast Name:%s\nPhone:%d\nEmail:%s\nPosition:%s\n\n",
		contacts[id-1].ID,
		contacts[id-1].FirstName,
		contacts[id-1].LastName,
		contacts[id-1].Phone,
		contacts[id-1].Email,
		contacts[id-1].Position,
	)
}

func (*Contact) Create() {

	var contact Contact

	size := uint8(len(contacts))

	contact.ID = size + 1

	fmt.Printf("Input First Name: ")
	fmt.Scanf("%s", &contact.FirstName)
	fmt.Printf("Input Last Name: ")
	fmt.Scanf("%s", &contact.LastName)
	fmt.Printf("Input Phone: ")
	fmt.Scanf("%d", &contact.Phone)
	fmt.Printf("Input Email: ")
	fmt.Scanf("%s", &contact.Email)
	fmt.Println("Input Position: ")
	fmt.Scanf("%s", &contact.Position)

	contacts = append(contacts, contact)
}

func (*Contact) Update() {

	var id uint8

	fmt.Println("Enter contact ID:")
	fmt.Scan(&id)

	for {
		var choice int
		fmt.Printf("\n1. First Name\n2. Last Name\n3. Phone\n4. Email\n5. Position\n0. Stop\n\n")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			{
				fmt.Printf("Input First Name: ")
				fmt.Scanf("%s", &contacts[id-1].FirstName)
			}
		case 2:
			{
				fmt.Printf("Input Last Name: ")
				fmt.Scanf("%s", &contacts[id-1].LastName)
			}
		case 3:
			{
				fmt.Printf("Input Phone: ")
				fmt.Scanf("%d", &contacts[id-1].Phone)
			}
		case 4:
			{
				fmt.Printf("Input Email: ")
				fmt.Scanf("%s", &contacts[id-1].Email)
			}
		case 5:
			{
				fmt.Println("Input Position: ")
				fmt.Scanf("%s", &contacts[id-1].Position)
			}
		default:
			return
		}
	}
}

func (*Contact) Delete() {

	var id uint8

	fmt.Println("Enter contact ID:")
	fmt.Scan(&id)

	for _, contact := range contacts {

		if id == contact.ID {
			id--
			copy(contacts[id:], contacts[id+1:])
			contacts[len(contacts)-1] = Contact{}
			contacts = contacts[:len(contacts)-1]
			return
		}
	}
}

func Contacts(number int) {

	var contact Contact

	switch number {
	case 1:
		contact.GetAll()
	case 2:
		contact.Get()
	case 3:
		contact.Create()
	case 4:
		contact.Update()
	case 5:
		contact.Delete()
	default:
		os.Exit(1)
	}

}
