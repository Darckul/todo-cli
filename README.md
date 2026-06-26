# todo-cli

A command-line task manager written in Go. Tasks are stored locally in a JSON file.

## Features

- Add tasks
- List all tasks with status
- Mark tasks as done
- Delete tasks
- Persistent storage via `tasks.json`

## Requirements

- Go 1.21+

## Installation

```bash
git clone https://github.com/Darckul/todo-cli.git
cd todo-cli
go build -o todo .
```

## Usage

```bash
# Add a task
go run . add "Buy groceries"

# List all tasks
go run . list

# Mark task #1 as done
go run . done 1

# Delete task #1
go run . delete 1
```

## Example

```
$ go run . add "Buy groceries"
Добавлено [1]: Buy groceries

$ go run . add "Read a book"
Добавлено [2]: Read a book

$ go run . list
[ ] [1] Buy groceries
[ ] [2] Read a book

$ go run . done 1
Выполнено [1]: Buy groceries

$ go run . list
[x] [1] Buy groceries
[ ] [2] Read a book
```

## Project Structure

```
todo-cli/
├── main.go     # entry point, command routing
├── todo.go     # Task struct, JSON load/save
└── tasks.json  # runtime data (not tracked by git)
```
