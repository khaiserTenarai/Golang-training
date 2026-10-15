package repository

import "gorm.io/gorm"

// ErrNotFound is what every repository returns when a row does not exist.
var ErrNotFound = gorm.ErrRecordNotFound
