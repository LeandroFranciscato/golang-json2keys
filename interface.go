package json2keys

type IJson2Keys interface {
	Parse(jsonStr string) (keys map[string]interface{}, err error)
}
