package chainOfResponsibilities

import (
	"bytes"
	"strings"
)

type Handler interface {
	Handle(string) string
	SetNextHandler(Handler)
}

//----------大写转换器-----------//
type UpperFilterHandler struct {
	nextHandler Handler
}
func (handle *UpperFilterHandler) Handle(rawString string)(string){
	newString := strings.ToUpper(rawString)
	if handle.nextHandler != nil{
		return handle.nextHandler.Handle(rawString)
	}
	return newString
}
func (handle *UpperFilterHandler) SetNextHandler(nextHandler Handler) {
	handle.nextHandler = nextHandler
}
//----------大写转换器-----------//

//----------多字节过滤器-----------//
type MultiByteFilterHandle struct {
	nextHandler Handler
}
func (handle *MultiByteFilterHandle) Handle(rawString string)(string){
	rawByte := []byte(rawString)
	rawRune := bytes.Runes(rawByte)
	ripeByte := make([]byte, len(rawString))
	for _, value := range rawRune{
		if len(string(value)) == 1{
			ripeByte = append(ripeByte, byte(value))
		}
	}
	if handle.nextHandler !=nil{
		return handle.nextHandler.Handle(string(ripeByte))
	}
	return string(ripeByte)
}
func (handle *MultiByteFilterHandle) SetNextHandler(nextHandler Handler) {
	handle.nextHandler = nextHandler
}
//----------多字节过滤器-----------//