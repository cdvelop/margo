package model

type Field struct {
	Name     string //id_usuario, apellido, addres, city etc
	DataType string //segun base de datos string,int,bool
	Legend   string //como se mostrara al usuario el campo en el encabezado ej: "name" por "Nombre Usuario"
	Input    Type   //text,radio,date,checkbox,rut,patente, file,password,tel,hidden,button etc...
	Validate bool   //si se requiere ser validado o no
}
