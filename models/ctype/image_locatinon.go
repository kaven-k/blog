package ctype

import "encoding/json"

type ImageLocation int

const (
	Local ImageLocation = 1 // 本地
	QiNiu ImageLocation = 2 // 七牛
)

func (s ImageLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageLocation) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case QiNiu:
		str = "七牛云"
	default:
		str = "其他"
	}
	return str
}
