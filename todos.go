package main

	
import (
    "fmt"
    "os"
    "bufio"
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

func show_done_task_list(){
    file,err := os.Open("todos.json")

    if(err != nil){
        fmt.Println("error = ",err)
    }

    data:= bufio.NewScanner(file)


    var lines[]string

    for data.Scan(){
        lines = append(lines, data.Text())
    }
    file.Close()

    for index,line:= range lines{
        fmt.Println(index," ",line)
    }
}

func main(){
	// help()
    show_done_task_list()
}