package model 

import (
  "gorm.io/gorm"
) // import

var (
  db	*gorm.DB
) // var()

func Init(_db *gorm.DB) {
    db = _db
} // Init()