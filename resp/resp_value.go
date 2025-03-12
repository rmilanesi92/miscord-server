package resp

import "strconv"

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

// Simplified RespValue int constructor
func NewArray(list []RespValue) RespValue {
    return RespValue{ Kind: ARRAY, Value: list}
}

// Convert the RespValue in a byte[] of its resp string representation
func (v *RespValue) ToBytes() []byte {
    switch(v.Kind) {
    case STR, ERR:
        return v.convertStr()
    case BULK_STR:
        return v.convertBulkStr()
    default:
        return []byte{}
    }
}

// Convert a str RespValue in byte[]
// In case of error an empty byte array is returned
func (v *RespValue) convertStr() []byte {
    var bytes []byte
    realValue, ok := v.Value.(string)
    if !ok {
        return []byte{}
    }
    bytes = append(bytes, v.Kind)
    bytes = append(bytes, realValue...)
    bytes = append(bytes, '\r', '\n')
    return bytes
}

// Convert a bulk str RespValue in byte[]
// In case of error an empty byte array is returned
func (v *RespValue) convertBulkStr() []byte {
    var bytes []byte
    realValue, ok := v.Value.(string)
    if !ok {
        return []byte{}
    }
    bytes = append(bytes, v.Kind)
    bytes = append(bytes, strconv.Itoa(len(realValue))...)
    bytes = append(bytes, '\r', '\n')
    bytes = append(bytes, realValue...)
    bytes = append(bytes, '\r', '\n')
    return bytes
}
