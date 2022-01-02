package main

import "errors"

var (
	ErrOpeningDatabase = errors.New("unable to open the database")
	ErrEmptyBucketName = errors.New("bucket name cannot be empty")
	ErrCreatingBucket  = errors.New("unable to create bucket")
	ErrInvalidType     = errors.New("invalid type, expected struct/map")
)
