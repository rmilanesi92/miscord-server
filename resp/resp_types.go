package resp

const (
    // Simple string that can't contains CRLF inside of it
    // +OK\r\n
    STR      = '+'

    // A bulk string represents a single binary string
    // $<length>\r\n<data>\r\n
    BULK_STR = '$'
    
    // Same as BULK_STR but providing data encoding information
    // =<length>\r\n<encoding>:<data>\r\n
    VERB_STR = '='

    // CRLF-terminated string representing a signed, base-10, 64-bit integer.
    // :[<+|->]<value>\r\n
    INT      = ':'

    // Similar to STR but threated as exceptions
    // -Error message\r\n
    ERR      = '-'

    // Same purpose of ERR but with the functionality of BULK_STR
    // !<length>\r\n<error>\r\n
    BULK_ERR = '!'

    // Indicate a collection of elements
    // *<number-of-elements>\r\n<element-1>...<element-n>
    ARRAY    = '*'

    // The null data type represents non-existent values.
    // _\r\n
    NULL     = '_'

    // Represent true or false values
    // #<t|f>\r\n
    BOOL     = '#'

    // The Double RESP type encodes a double-precision floating point value
    // ,[<+|->]<integral>[.<fractional>][<E|e>[sign]<exponent>]\r\n
    DOUBLE   = ','

    // This type encode integer outside the range of signed 64-bit integers.
    // ([+|-]<number>\r\n
    BIG_NUM  = '('
    
    // Collection of key-value tuples.
    // %<number-of-entries>\r\n<key-1><value-1>...<key-n><value-n>
    MAP      = '%'

    // Like MAP but describes only "metadata" of the response, never real data
    // |1\r\n
    //     +key-popularity\r\n
    //     %2\r\n
    //         $1\r\n
    //         a\r\n
    //         ,0.1923\r\n
    //         $1\r\n
    //         b\r\n
    //         ,0.0012\r\n
    // *2\r\n
    //    :2039123\r\n
    //    :9543892\r\n
    ATTR     = '|'

    // Unordered array with unique items
    // ~<number-of-elements>\r\n<element-1>...<element-n>
    SET      = '~'

    // RESP's pushes contain out-of-band data
    // ><number-of-elements>\r\n<element-1>...<element-n>
    PUSH     = '>'
)
