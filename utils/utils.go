package utils

import "strings"

const (
	alphabet     = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetsize = uint64(len(alphabet))
)

func Encode(number uint64) string {
	var builder strings.Builder
	builder.Grow(IdLength)
	var values [IdLength]byte

	for i := 0; i < IdLength; i++ {
		values[i] = alphabet[((number + uint64(i)) % alphabetsize)]
		number /= alphabetsize
	}
	for i := IdLength - 1; i >= 0; i-- {
		builder.WriteByte(values[i])
	}

	return builder.String()
}

func GenerateUniqueId(intKey uint64) string {
	return Encode((intKey * MagicPrime) & KeyMask)
}
