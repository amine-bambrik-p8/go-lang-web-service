package models

// GetViewModel will be called before the model is sent for modifying the content of the model
type Model interface {
	GetViewModel() interface{}
}
