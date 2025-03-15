package command

import (
	"fmt"
	"strings"

	"github.com/rmilanesi92/miscord-server/resp"
)

// Command Representation
type Command struct {
    Name string
    Exec func([]resp.RespValue) resp.RespValue
}

// Will hold the command list
var CommandList = make(map[string]Command)

// Handle command parsing and execution
func Handle(input resp.RespValue) resp.RespValue {
    if input.Kind != resp.ARRAY {
        return resp.NewErrorFromMsg("Array format expected")
    } 

    parsedInput, ok := input.Value.([]resp.RespValue)
    
    if !ok || len(parsedInput) == 0 {
        return resp.NewErrorFromMsg("Command format issue")
    }
   
    commandName := parsedInput[0].Value.(string)
    commandName = strings.ToUpper(commandName)

    args := parsedInput[1:]

    command, ok := CommandList[commandName]
    if !ok {
        return resp.NewErrorFromMsg(
            fmt.Sprintf("Invalid command: %s", commandName),
        )
    }

    return command.Exec(args)
}

// Register a Command instance in CommandList map
func RegisterCommand(cmd Command) {
    CommandList[strings.ToUpper(cmd.Name)] = cmd
}

// Register a List of Command instance in CommandList map
func RegisterCommandList(list []Command) {
    for _, cmd := range list {
        CommandList[strings.ToUpper(cmd.Name)] = cmd
    }
}
