package utils

const (
	BaseKeyName  = "key_index"
	IdLength     = 7
	BinaryLength = 41
	MagicPrime   = uint64(70936234049)
	KeyMask      = (uint64(1) << BinaryLength) - 1
)
