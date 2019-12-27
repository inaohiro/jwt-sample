package codes

import "errors"

//type NotFound struct{}
//
//func (e *NotFound) Error() string {
//	return "no such entity"
//}

var (
	NotFound      = errors.New("no such entity")
	InvalidEntity = errors.New("invalid entity")
)
