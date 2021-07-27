package err

type ErrorCode int

//go:generate stringer -type=ErrorCode
const (
	ERR_INVALID_PARAM ErrorCode = iota + 3000
	ERR_INVALID_EMAIL
)

type Error struct {
	code int
}