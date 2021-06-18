package model

//adaptador html input
type Type interface {
	Input() string        //representacion en html
	Validate(string) bool //como sera validado
}
