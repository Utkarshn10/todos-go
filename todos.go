package main

	
import (
    "fmt"
)

func help(){
	output_str:= `
	Use following commands to perform actions:

    1. List all todos                   go run todos.go ls
    2. Add a new todo                   go run todos.go add "Todo_Name"
    3. Delete a todo                    go run todos.go rm "Todo_Index"
    3. Update a todo                    go run todos.go update "Todo_Index" new_todo
    4. Mark a todo as done              go run todos.go done "Todo_Index"
    5. Show all the completed todos     go run todos.go ls_done "Todo_Index"
	`

	fmt.Println(output_str)
}

func main(){
	help()
}