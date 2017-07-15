package utils


import (
    "fmt"
    "math/rand"
    crypto "crypto/rand"
)

type Identifier string

func (i Identifier) ToString() string {
    return fmt.Sprintf(
        "%X-%X-%X-%X-%X",
        i[0:4],
        i[4:6],
        i[6:8],
        i[8:10],
        i[10:],
    )
}

func RandomId() Identifier {
    val := make([]byte, 16)
    rand.Read(val)
    return Identifier(val)
}


func SecureId() Identifier {
    val := make([]byte, 16)
    crypto.Read(val)
    return Identifier(val)
}

