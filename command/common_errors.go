package command

import (
	"fmt"
	"github.com/rmilanesi92/miscord-server/resp"
)

// InvalidArgumentErr common error for many commands
func InvalidArgumentErr(commandName string) resp.RespValue {
    return resp.NewErrorFromMsg(
        fmt.Sprintf(
            "Invalid number of arguments for command '%s'", 
            commandName,
        ),
    )
}
