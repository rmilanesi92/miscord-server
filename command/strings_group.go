package command

import (
	"github.com/rmilanesi92/miscord-server/resp"
    "github.com/rmilanesi92/miscord-server/data"
)

// Get the value of key. 
// If the key does not exist the special value nil is returned. 
// An error is returned if the value stored at key is not a string, 
func get(args []resp.RespValue) resp.RespValue {
    if len(args) != 1 {
        return InvalidArgumentErr("get")
    }

    key, ok := args[0].Value.(string)
    if !ok {
        return resp.NewErrorFromMsg("Retrieved value is not a string")
    }

    data.DBSetMutex.RLock()
    value, ok := data.DBSet[key]
    data.DBSetMutex.RUnlock()

    if !ok {
        return resp.Null()
    }

    return resp.NewBulkString(value)
}

// Set key to hold the string value. 
// If key already holds a value, it is overwritten, regardless of its type. 
func set(args []resp.RespValue) resp.RespValue {
   if len(args) != 2 {
        return InvalidArgumentErr("set")
   } 

   key := args[0].Value.(string)
   value := args[1].Value.(string)

   data.DBSetMutex.Lock()
   data.DBSet[key] = value
   data.DBSetMutex.Unlock()
    
   return resp.NewString("Ok")
}

// Init function used to add commands to the common list
func init() {
    RegisterCommandList([]Command {
        {
            Name: "get",
            Exec: get,
        },
        {
            Name: "set",
            Exec: set,
        },
    })
}
