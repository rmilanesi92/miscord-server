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
        {
            name: "Error",
            input: RespValue{
                Kind: ERR,
                Value: "ERR test example",
            },
            expected: "-ERR test example\r\n",
        },
        {
            name: "Bulk String",
            input: RespValue{
                Kind: BULK_STR,
                Value: "A\r\nBulk\r\nString",
            },
            expected: "$15\r\nA\r\nBulk\r\nString\r\n",
        },
        {
            name: "Array of simple Strings",
            input: RespValue{
                Kind: ARRAY,
                Value: []RespValue{
                    { Kind: STR, Value: "uno"},
                    { Kind: STR, Value: "dos"},
                },
            },
            expected: "*2\r\n+uno\r\n+dos\r\n",
        },
        {
            name: "Array of bulk Strings",
            input: RespValue{
                Kind: ARRAY,
                Value: []RespValue{
                    { Kind: BULK_STR, Value: "uno"},
                    { Kind: BULK_STR, Value: "dos"},
                },
            },
            expected: "*2\r\n$3\r\nuno\r\n$3\r\ndos\r\n",
        },
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            // Execute the method
            bytes := tc.input.ToBytes()

            actual := string(bytes)
            
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
