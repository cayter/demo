package err

type ErrorCode int

//go:generate stringer -type=ErrorCode
//go:generate go run ../generator/error.go -type=ErrorCode
const (
	ERR_INVALID_PARAM ErrorCode = iota + 3000
	ERR_INVALID_EMAIL
)

const (
	ERR_INVALID_USERNAME ErrorCode = iota + 4000
)
