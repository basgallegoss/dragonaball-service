package api

import "errors"

var (
	ErrCharacterNotFound = errors.New("character not found in Dragonball API")
	ErrDecodeResponse    = errors.New("failed to decode Dragonball API response")
	ErrInvalidURL        = errors.New("invalid URL for Dragonball API")
	ErrRequestFailed     = errors.New("request to Dragonball API failed")
	ErrUnexpectedStatus  = errors.New("Dragonball API returned unexpected status")
)
