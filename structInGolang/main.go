package main

import (
	"fmt"
	"os"
	"time"

	c "app/contacts"
	t "app/tasks"
)

func main() {

	fmt.Println("CRUD app")

	var choice int
	fmt.Printf("1. Contact list\n2. Taks list\n0. Exit\n\n")
	fmt.Scan(&choice)

	for {
		switch choice {
		case 1:
			{
				var contactNumber int
				fmt.Printf("\n1. Get All\n2. Get\n3. Create\n4. Update\n5. Delete\n6. Menu\n0. Exit\n\n")
				fmt.Scan(&contactNumber)

				if contactNumber == 6 {
					break
				}
				c.Contacts(contactNumber)
			}
		case 2:
			{
				var taskNumber int
				fmt.Printf("\n1. Get All\n2. Get\n3. Create\n4. Update\n5. Delete\n6. Menu\n0. Exit\n\n")
				fmt.Scan(&taskNumber)

				if taskNumber == 6 {
					break
				}

				t.Tasks(taskNumber)
			}
		default:
			os.Exit(1)
		}

		time.Sleep(time.Second)
	}
}
