package handlers

import "fmt"

const apiDateFormat = "2006-Jan-02"

var ErrCreateHandlerWithNilMapper = fmt.Errorf("cannot create handler with nil mapper")
