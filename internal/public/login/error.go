package login

const (
	ERR = 1
)

func ErrStatusText(code int) string {
	switch code {
	case OK:
		return "ok"
	default:
		return ""
	}
}
