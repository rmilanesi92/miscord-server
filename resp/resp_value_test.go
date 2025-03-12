package resp

import "testing"


func Test_ToBytes(t *testing.T) {
    cases := []struct {
        name string
        input RespValue
        expected string
    } {
        {
            name: "Simple String",
            input: RespValue{
                Kind: STR,
                Value: "SimpleString",
            },
            expected: "+SimpleString\r\n",
        },
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            // Execute the method
            actual := tc.input.ToBytes()

            if string(actual) != tc.expected {
                t.Errorf(
                    "Unexpected result.\nExpected: %+v\nActual: %+v",
                    tc.expected,
                    actual,
                )
            }
        })
    }
}
