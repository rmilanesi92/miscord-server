package resp

import (
	"bytes"
	"reflect"
	"testing"
)

// Tests public method Read agains different RESP types
func TestRead(t *testing.T) {
    cases := []struct{
        name        string
        input       string
        expected    RespValue
    } {
        {
            name: "Simple String",
            input: "+test\r\n",
            expected: RespValue{ Kind: STR, Value: "test"},
        },
        {
            name: "Bulk String",
            input: "$6\r\nte\r\nst\r\n",
            expected: RespValue{ Kind: BULK_STR, Value: "te\r\nst"},
        },
        {
            name: "Positive Integer",
            input: ":+892\r\n",
            expected: RespValue{ Kind: INT, Value: 892},
        },
        {
            name: "Negative Integer",
            input: ":-144\r\n",
            expected: RespValue{ Kind: INT, Value: -144}, },
    }
     for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) { 
            // Create buffer string that act as input buffer
            input := bytes.NewBufferString(tc.input)

            // Instantiate the reader
            resp := NewRespReader(input)  

            // Perform the reading
            actual := resp.Read()

            // Compare the actual result with the expected
            if !reflect.DeepEqual(actual, tc.expected) {
                t.Errorf(
                    "Unexpected result.\nExpected: %+v\nActual: %+v",
                    tc.expected,
                    actual,
                )
            }
        })   
    }
}
