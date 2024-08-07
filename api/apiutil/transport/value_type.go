package transport

type CustomHeader string

const (
	CUSTOM_HEADERS CustomHeader = "CustomHeaders"
)

type Values struct {
	HeaderMap map[string]string
}

func (v Values) Get(key string) string {
	if v.HeaderMap == nil {
		return ""
	}
	return v.HeaderMap[key]
}
