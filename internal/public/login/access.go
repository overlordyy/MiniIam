package login

const (
	OK = 0
)

func AclStatusText(code int) string {
	switch code {
	case OK:
		return "ok"
	default:
		return ""
	}
}
