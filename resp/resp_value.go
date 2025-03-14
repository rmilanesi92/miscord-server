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

// Simplified RespValue int constructor
func Null() RespValue {
    return RespValue{ Kind: NULL}
}

// Convert the RespValue in a byte[] of its resp string representation
// In case of error an empty byte array is returned
func (v *RespValue) ToBytes() []byte {
    switch(v.Kind) {
    case STR, ERR:
        return v.convertStr()
    case BULK_STR:
        return v.convertBulkStr()
    case ARRAY:
        return v.convertArray()
    case NULL:
        return v.convertNull()
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

// Convert an array RespValue in byte[]
// In case of error an empty byte array is returned
func (v *RespValue) convertArray() []byte {
    var bytes []byte
    realValue, ok := v.Value.([]RespValue)
    if !ok {
        return []byte{}
    }
    bytes = append(bytes, v.Kind)
    bytes = append(bytes, strconv.Itoa(len(realValue))...)
    bytes = append(bytes, '\r', '\n')
    for _, value := range realValue {
        bytes = append(bytes, value.ToBytes()...)
    }
    return bytes
}

// Convert a Null RespValue in byte[]
func (v *RespValue) convertNull() []byte {
    var bytes []byte
    bytes = append(bytes, v.Kind)
    bytes = append(bytes, '\r', '\n')
    return bytes
}


