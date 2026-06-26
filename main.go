package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: todo <команда> [аргументы]")
		fmt.Println("Команды: add, list, done, delete")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		cmdAdd()
	case "list":
		cmdList()
	case "done":
		cmdDone()
	case "delete":
		cmdDelete()
	default:
		fmt.Printf("Неизвестная команда: %s\n", command)
		os.Exit(1)
	}
}

func cmdAdd() {
	if len(os.Args) < 3 {
		fmt.Println("Использование: todo add <текст задачи>")
		os.Exit(1)
	}

	text := os.Args[2]

	tasks, err := load()
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		os.Exit(1)
	}

	newTask := Task{
		ID:   nextID(tasks),
		Text: text,
		Done: false,
	}
	tasks = append(tasks, newTask)

	if err := save(tasks); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		os.Exit(1)
	}

	fmt.Printf("Добавлено [%d]: %s\n", newTask.ID, newTask.Text)
}


func nextID(tasks []Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}

func cmdList() {
	tasks, err := load()
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		os.Exit(1)
	}

	if len(tasks) == 0 {
		fmt.Println("Задач ещё нет")
		return
	}

	for _, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%s [%d] %s\n", status, t.ID, t.Text)
	}
}

func cmdDone() {
	if len(os.Args) < 3 {
		fmt.Println("Использование: todo done <id>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("ID должен быть числом")
		os.Exit(1)
	}

	tasks, err := load()
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		os.Exit(1)
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true 
			if err := save(tasks); err != nil {
				fmt.Println("Ошибка сохранения:", err)
				os.Exit(1)
			}
			fmt.Printf("Выполнено [%d]: %s\n", t.ID, t.Text)
			return
		}
	}

	fmt.Printf("Задача #%d не найдена\n", id)
	os.Exit(1)
}

func cmdDelete() {
	if len(os.Args) < 3 {
		fmt.Println("Использование: todo delete <id>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("ID должен быть числом")
		os.Exit(1)
	}

	tasks, err := load()
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		os.Exit(1)
	}

	for i, t := range tasks {
		if t.ID == id {
			
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err := save(tasks); err != nil {
				fmt.Println("Ошибка сохранения:", err)
				os.Exit(1)
			}
			fmt.Printf("Удалено [%d]: %s\n", t.ID, t.Text)
			return
		}
	}

	fmt.Printf("Задача #%d не найдена\n", id)
	os.Exit(1)
}
