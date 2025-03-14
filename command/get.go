package command

import (
	"github.com/rmilanesi92/miscord-server/resp"
    "github.com/rmilanesi92/miscord-server/data"
)

func get(args []resp.RespValue) resp.RespValue {
    if len(args) != 1 {
        return resp.NewErrorFromMsg("Wrong number of argument for GET")
    }

    key := args[0].Value.(string)

    data.DBSetMutex.RLock()
    value, ok := data.DBSet[key]
    data.DBSetMutex.RUnlock()

    if !ok {
        return resp.Null()
    }

    return resp.NewBulkString(value)
}

// Init function used to add the command to the common list
func init() {
    RegisterCommand(Command{
        Name: "Get",
        Exec: get,
    })
}
