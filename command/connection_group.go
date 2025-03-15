package command

import "github.com/rmilanesi92/miscord-server/resp"


// Returns PONG if no argument is provided, 
// otherwise return a copy of the argument as a bulk. 
// This command is useful for:
// 1. Testing whether a connection is still alive.
// 2. Verifying the server's ability to serve data
//    an error is returned when this isn't the case 
//    (e.g., during load from persistence or accessing a stale replica).
// 3. Measuring latency.
func ping(args []resp.RespValue) resp.RespValue {
    if len(args) > 1 {
        return InvalidArgumentErr("ping")
    }
    
    pong := resp.NewBulkString("PONG")
    if len(args) == 0 {
        return pong
    }

    value, ok := args[0].Value.(string)
    if !ok {
       return pong 
    }

    return resp.NewBulkString(value)
}

func init() {
    RegisterCommandList([]Command{
        {
            Name: "ping",
            Exec: ping,
        },
    })    
}
