package dberrors

import "errors"

 var (
	ErrIndexNotInFields = errors.New("Index not in Field!")
	ErrIndexType = errors.New("You can only choose between string|int")
	ErrFailedToconnect = errors.New("Failed to connect to the database!") 
	ErrDbName = errors.New("Database doesn't exist!")
	ErrDbInsert = errors.New("Unable to add Database!, Check the fields")
)