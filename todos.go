package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
    "bufio"
    "strings"
    "strconv"
	"os"
)

func help(){
	output_str:= `

                WELCOME TO TODOS
    ----------------------------------------------------------------

	Here is a list of commands you can try:

    1. List all todos                   ls
    2. Add a new todo                   add "Todo_Name"
    3. Delete a todo                    rm "Todo_Index"
    3. Update a todo                    update "Todo_Index" new_todo
    4. Mark a todo as done              done "Todo_Index"
    5. Show all the completed todos     ls_done "Todo_Index"
	`

	fmt.Println(output_str)
}

func readFile() ([]byte,error){
    file,err := os.Open("todos.json")

    if(err != nil){
        fmt.Println("error = ",err)
    }

    data,error := ioutil.ReadAll(file)
    file.Close()

    if error != nil {
        fmt.Println("Error:", error)
        return nil, error // Return the error
    }

    return data,nil
}

func writeToFile(updatedData []byte) {
    if err := ioutil.WriteFile("todos.json", updatedData, 0644); err != nil {
        fmt.Println("Error writing file:", err)
        return 
    }
    return
}

func addTask(task string){
    data,error := readFile()

    if error != nil {
        fmt.Println("Error reading file:", error)
        return
    }

    var contentMap map[string]interface{}
    json.Unmarshal([]byte(data),&contentMap)

    key := strconv.Itoa(len(contentMap) + 1)
    contentMap[key] = task

    updatedData,err := json.Marshal(contentMap)

    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }

    writeToFile(updatedData)

    fmt.Println("Task Added")
}

func deleteTask(taskKey string){
    data,error := readFile()

    if error != nil {
        fmt.Println("Error reading file:", error)
        return
    }

    var contentMap map[string]interface{}
    json.Unmarshal([]byte(data),&contentMap)
    fmt.Println(contentMap," ",taskKey)
    delete(contentMap,taskKey)
    newContentMap := make(map[string]interface{})

    index :=1 

    for _,value := range(contentMap) {
        newKey := strconv.Itoa(index)
        newContentMap[newKey] = value
        index++
    }

    updatedData,err := json.Marshal(newContentMap)

    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }

    writeToFile(updatedData)
    fmt.Println("Task Deleted")
}

func updateTask(task string){
    data,error := readFile()

    if error != nil {
        fmt.Println("Error reading file:", error)
        return
    }

    var contentMap map[string]interface{}
    json.Unmarshal([]byte(data),&contentMap)

}

func completedTasks(task string){}

func markTaskAsDone(task string){}

func showActiveTaskList(){
    file,err := os.Open("todos.json")

    if(err != nil){
        fmt.Println("error = ",err)
    }

    data,_ := ioutil.ReadAll(file)
    file.Close()

    var contentMap map[string]interface{}
    json.Unmarshal([]byte(data),&contentMap)

    // output content
    fmt.Println("--------------------------------")
    fmt.Println("ACTIVE TASKS")
    fmt.Println("--------------------------------")
    for key,value := range contentMap{
        fmt.Println(key,value)
    }
}

func main(){
    reader := bufio.NewReader(os.Stdin)

	help()

    fmt.Println("Enter command to continue....")
    fmt.Println()

    userInput, _ := reader.ReadString('\n')
    userInput = strings.TrimSpace(userInput)

    actions := strings.SplitN(userInput," ",2)

    
    switch actions[0]{
        case "add":
            task := actions[1]
            addTask(task)
        case "ls":
            showActiveTaskList()
        case "rm":
            task := actions[1]
            deleteTask(task)
        case "update":
            task := actions[1]
            updateTask(task)
        case "ls_done":
            task := actions[1]
            completedTasks(task)
        case "done":
            task := actions[1]
            markTaskAsDone(task) 
        default:
            fmt.Println("Enter command to continue....")
    }
}