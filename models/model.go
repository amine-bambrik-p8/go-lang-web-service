package models

// Use for hidding field before sending to the client
type Model interface {
	GetViewModel() interface{}
}
