package controllers

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "os/exec"
)

func CommandManaer(c *gin.Context){
    commandName := c.PostForm("command")
    commandLine := c.PostForm("method")
    fmt.Println("comand :", commandName)
    fmt.Println("comand line:", commandLine)

    var cmd *exec.Cmd
    
    if commandLine == "cmd"{
        cmd = exec.Command( "cmd", "/C", commandName)
    }else if commandLine == "terminal"{
        cmd = exec.Command("wsl", "bash", "-c", commandName)
    }
    
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("error:",err)
        c.JSON(400, gin.H{
		    "error" : output,
	    })
    } else{
        fmt.Println("output:",output)
        c.String(200, string(output))
    }
    
}

// func main() {
//     // Command to run (for example: 'ping google.com -n 1' on Windows)
//     // cmd := exec.Command("ping", "google.com", "-n", "1") // "-c" for Unix, use "-n" for Windows

// 	cmd := exec.Command( "cmd", "/C","cd kejrvn")

//     // Run it and capture the output
//     output, err := cmd.CombinedOutput()

//     if err != nil {
//         fmt.Println("Error:", err)
//     }

//     fmt.Println(string(output))
// }