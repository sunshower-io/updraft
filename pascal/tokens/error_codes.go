package tokens


type ErrorCode string



const (
    ALREADY_FORWARDED           ErrorCode = "Already specified in FORWARD"
    IDENTIFIER_REDEFINED        ErrorCode = "Redefined identifier"
    IDENTIFIER_UNDEFINED        ErrorCode = "Undefined identifier"
    INCOMPATIBLE_ASSIGNMENT     ErrorCode = "Incompatible types"
    INVALID_ASSIGNMENT          ErrorCode = "Invalid assignment statement"
    INVALID_CHARACTER           ErrorCode = "Invalid character"
    INVALID_CONSTANT            ErrorCode = "Invalid constant"
    INVALID_EXPONENT            ErrorCode = "Invalid exponent"
    INVALID_EXPRESSION          ErrorCode = "Invalid expression"
    INVALID_FIELD               ErrorCode = "Invalid field"
    INVALID_FRACTION            ErrorCode = "Invalid fraction"
    INVALID_IDENTIFIER_USAGE    ErrorCode = "Invalid identifier usage"
    INVALID_INDEX_TYPE          ErrorCode = "Invalid index type"
    
    INVALID_NUMBER              ErrorCode = "Invalid Number"
    INVALID_STATEMENT           ErrorCode = "Invalid statement"
    
    INVALID_SUBRANGE_TYPE       ErrorCode = "Invalid subrange type"
    
    INVALID_TARGET              ErrorCode = "Invalid assignment target"
    INVALID_TYPE                ErrorCode = "Invalid type"
    INVALID_VAR_PARAM           ErrorCode = "Invalid VAR parameter"
    INVALID_GT_MAX              ErrorCode = "Minimum limit greater than maximum limit"
    MISSING_BEGIN               ErrorCode = "Missing BEGIN"
    MISSING_COLON               ErrorCode = "Missing ':'"
    MISSING_COLON_EQUALS        ErrorCode = "Missing ':='"
    MISSING_COMMA               ErrorCode = "Missing ','"
    MISSING_CONSTANT            ErrorCode = "Missing constant"
    MISSING_DO                  ErrorCode = "Missing 'DO'"
    MISSING_DOT_DOT             ErrorCode = "Missing '..'"
    MISSING_END                 ErrorCode = "Missing END"
    MISSING_EQUALS              ErrorCode = "Missing Equals"
    MISSING_FOR_CONTROL         ErrorCode = "Missing for control variable"
    MISSING_IDENTIFIER          ErrorCode = "Missing identifier"
    MISSING_LEFT_BRACKET        ErrorCode = "Missing '['"
    MISSING_OF                  ErrorCode = "Missing OF"
    MISSING_PERIOD              ErrorCode = "Missing '.'"
    MISSING_PROGRAM             ErrorCode = "Missing 'PROGRAM'"
    MISSING_RIGHT_BRACKET       ErrorCode = "Missing ']'"
    MISSING_RIGHT_PAREN         ErrorCode = "Missing ')'"
    MISSING_SEMICOLON           ErrorCode = "Missing ';'"
    MISSING_THEN                ErrorCode = "Missing THEN"
    MISSING_TO_DOWNTO           ErrorCode = "Missing 'TO' or 'DOWNTO'"
    MISSING_UNTIL               ErrorCode = "Missing 'Until'"
    MISSING_VARIABLE            ErrorCode = "Missing variable"
    CASE_CONSTANT_REUSED        ErrorCode = "CASE constant reused"
    NOT_CONSTANT_IDENTIFIER     ErrorCode = "Not a constant identifier"
    NOT_RECORD_VARIABLE         ErrorCode = "Not a record variable"
    NOT_TYPE_IDENTIFIER         ErrorCode = "Not a type identifier"
    RANGE_INTEGER               ErrorCode = "Integer literal out of range"
    RANGE_REAL                  ErrorCode = "Real literal out of range"
    STACK_OVERFLOW              ErrorCode = "Stack overflow"
    TOO_MANY_LEVELS             ErrorCode = "Nesting level too deep"
    TOO_MANY_SUBSCRIPTS         ErrorCode = "Too many subscripts"
    UNEXPECTED_EOF              ErrorCode = "Unexpected EOF"
    UNEXPECTED_TOKEN            ErrorCode = "Unexpected token"
    UNIMPLEMENTED               ErrorCode = "Unimplemented"
    UNRECOGNIZABLE              ErrorCode = "Unrecognizable input"
    WRONG_NUMBER_OF_PARAMS      ErrorCode = "Wrong number of parameters"
    
   
)
