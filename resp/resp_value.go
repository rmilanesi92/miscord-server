package resp

// Represent a RESP value
type RespValue struct {
    Kind byte
    Value interface{}
}

// Simplified RespValue error constructor from error interface
func NewError(err error) RespValue {
    return RespValue{ Kind: ERR, Value: err.Error()}
}

// Simplified RespValue error constructor from error message 
func NewErrorFromMsg(msg string) RespValue {
    return RespValue{ Kind: ERR, Value: msg}
}

// Simplified RespValue string constructor
func NewString(msg string) RespValue {
    return RespValue{ Kind: STR, Value: msg}
}

// Simplified RespValue string constructor
func NewBulkString(msg string) RespValue {
    return RespValue{ Kind: BULK_STR, Value: msg}
}

// Simplified RespValue int constructor
func NewInteger(num int) RespValue {
    return RespValue{ Kind: INT, Value: num}
}
