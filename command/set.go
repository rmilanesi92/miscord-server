package command

import (
	"github.com/rmilanesi92/miscord-server/resp"
    "github.com/rmilanesi92/miscord-server/data"
)

// Add a key value pair in the set DB
func set(args []resp.RespValue) resp.RespValue {
   if len(args) != 2 {
        return resp.NewErrorFromMsg("Wrong number of argument for SET")
   } 

   key := args[0].Value.(string)
   value := args[0].Value.(string)

   data.DBSetMutex.Lock()
   data.DBSet[key] = value
   data.DBSetMutex.Unlock()
    
   return resp.NewString("Ok")
}


// Init function used to add the command to the common list
func init() {
    RegisterCommand(Command{
        Name: "Set",
        Exec: set,
    })
}
