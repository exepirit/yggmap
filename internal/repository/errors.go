package repository

import "errors"

var ErrNoObjects = errors.New("requested object(s) not found in repository")
