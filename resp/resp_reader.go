package resp

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// A reader for RESP protocol compatible message
type RespReader struct {
    reader *bufio.Reader
}

// Resp Reader constructor function
func NewRespReader(reader io.Reader) *RespReader {
    return &RespReader{reader: bufio.NewReader(reader)}
}

// readLine reads a single line from the reader and trims any CRLF char
func (self *RespReader) readLine() (string, error) {
    line, err := self.reader.ReadString('\n')
    if err != nil {
        return "", err
    }
    return strings.TrimSuffix(line, "\r\n"), nil
}
// readIntByte reads a line, converts it to an int, and returns it
func (self *RespReader) readIntByte() (int, error) {
    value, err := self.readLine()
    if err != nil {
        return 0, err
    }

    i, err := strconv.ParseInt(value, 10, 0)
    if err != nil {
        return 0, err
    }
    return int(i), nil
}

// Public Read method that translate given input buffer into RespValue
func (self *RespReader) Read() RespValue {
    kind, err := self.reader.ReadByte()
    if err != nil {
        return NewError(err)
    }
    switch kind {
    case STR:
        return self.readString()
    case BULK_STR:
        return self.readBulkString()
    default:
        return NewErrorFromMsg("ERR unsupported type: " + string(kind))
    }
}

// Read simple string from buffer
func (self *RespReader) readString() RespValue {
    value, err := self.readLine()
    if err != nil {
        return NewErrorFromMsg("ERR reading string")
    }
    return NewString(value)
}

// Read bulk string from buffer
func (self *RespReader) readBulkString() RespValue {
    length, err := self.readIntByte()
    if err != nil {
        return NewError(err)
    }

    value := make([]byte, length)
    nWritten, err := io.ReadFull(self.reader, value)
    if err != nil {
        return NewError(err)
    }

    if nWritten != length {
        return NewErrorFromMsg("ERR Invalid format for bulk string")
    }

    return NewBulkString(string(value))
}
