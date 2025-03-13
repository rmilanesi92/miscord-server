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
var CommandList map[string]Command

// Handle command parsing and execution
func Handle(input resp.RespValue) resp.RespValue {
    if input.Kind != resp.ARRAY {
        return resp.NewErrorFromMsg("ERR array format expected")
    } 

    parsedInput, ok := input.Value.([]resp.RespValue)
    
    if !ok || len(parsedInput) == 0 {
        return resp.NewErrorFromMsg("ERR command format issue")
    }
   
    commandName := parsedInput[0].Value.(string)
    commandName = strings.ToUpper(commandName)

    args := parsedInput[1:]

    command, ok := CommandList[commandName]
    if !ok {
        return resp.NewErrorFromMsg(
            fmt.Sprintf("ERR invalid command: %s", commandName),
        )
    }

    return command.Exec(args)
}
