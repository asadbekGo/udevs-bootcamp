package tasks

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID        uint8
	Name      string
	Status    string
	Priority  string
	CreatedAt string
	CreatedBy string
	DueDate   string
}

var tasks []Task

func (*Task) GetAll() {

	for _, task := range tasks {

		fmt.Printf("\nId:%d\nName:%s\nStatus:%s\nPriority:%s\nCreatedAt:%s\nCreatedBy:%s\nDueDate:%s\n",
			task.ID,
			task.Name,
			task.Status,
			task.Priority,
			task.CreatedAt,
			task.CreatedBy,
			task.DueDate,
		)
	}
}

func (*Task) Get() {

	var id uint8

	fmt.Println("Enter task ID:")
	fmt.Scan(&id)

	fmt.Printf("\nId:%d\nName:%s\nStatus:%s\nPriority:%s\nCreatedAt:%s\nCreatedBy:%s\nDueDate:%s\n",
		tasks[id-1].ID,
		tasks[id-1].Name,
		tasks[id-1].Status,
		tasks[id-1].Priority,
		tasks[id-1].CreatedAt,
		tasks[id-1].CreatedBy,
		tasks[id-1].DueDate,
	)
}

func (*Task) Create() {

	var task Task

	size := uint8(len(tasks))

	task.ID = size + 1

	fmt.Printf("Input First Name: ")
	fmt.Scanf("%s", &task.Name)
	fmt.Printf("Input Status: ")
	fmt.Scanf("%s", &task.Status)
	fmt.Printf("Input Priority: ")
	fmt.Scanf("%s", &task.Priority)
	fmt.Println("Input CreatedBy: ")
	fmt.Scanf("%s", &task.CreatedBy)
	fmt.Printf("Input DueDate: ")
	fmt.Scanf("%s", &task.DueDate)

	task.CreatedAt = time.Now().Format(time.ANSIC)

	tasks = append(tasks, task)
}

func (*Task) Update() {

	var id uint8

	fmt.Println("Enter task ID:")
	fmt.Scan(&id)

	for {
		var choice int
		fmt.Printf("\n1. Name\n2. Status\n3. Priority\n4. DueData\n5. CreatedBy\n0. Stop\n\n")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			{
				fmt.Printf("Input Name: ")
				fmt.Scanf("%s", tasks[id-1].Name)
			}
		case 2:
			{
				fmt.Printf("Input Status: ")
				fmt.Scanf("%s", tasks[id-1].Status)
			}
		case 3:
			{
				fmt.Printf("Input Priority: ")
				fmt.Scanf("%s", tasks[id-1].Priority)
			}
		case 4:
			{
				fmt.Printf("Input DueDate: ")
				fmt.Scanf("%s", tasks[id-1].DueDate)
			}
		case 5:
			{
				fmt.Println("Input CreatedBy: ")
				fmt.Scanf("%s", tasks[id-1].CreatedBy)
			}
		default:
			return
		}
	}
}

func (*Task) Delete() {

	var id uint8

	fmt.Println("Enter task ID:")
	fmt.Scan(&id)

	for _, task := range tasks {

		if id == task.ID {
			id--
			copy(tasks[id:], tasks[id+1:])
			tasks[len(tasks)-1] = Task{}
			tasks = tasks[:len(tasks)-1]
			return
		}
	}
}

func Tasks(number int) {

	var task Task

	switch number {
	case 1:
		task.GetAll()
	case 2:
		task.Get()
	case 3:
		task.Create()
	case 4:
		task.Update()
	case 5:
		task.Delete()
	default:
		os.Exit(1)
	}

}
