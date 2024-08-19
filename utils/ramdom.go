package utils

import(
	"fmt"
	"math/rand"
	"strings"
	
)

const alphabet = "abcdefghijlmnopqrstuvxz"

func RandomStrig(number int)  string{
	var sb strings.Builder
	k := len(alphabet)

	for i := 0;i < number;i++{
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail()  string{
	return fmt.Sprintf("%s@email.com",RandomStrig(6))
}