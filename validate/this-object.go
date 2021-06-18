package validate

import (
	"fmt"
	"strings"

	"github.com/cdvelop/margo/model"
)

//This ..option: new,all
func This(objectIn Object, option string) (message string, ok bool) {
	var wrongFields []string
	modelTableObject := model.GetTableModelByName(objectIn.TableName())

	switch option {
	case "new":
		var fewer int
		// el modelo tiene primary key?
		if _, exist := (*modelTableObject)["id_"+objectIn.TableName()]; exist {
			fewer = 1
		}
		if len(objectIn.FieldValue()) != len((*modelTableObject))-fewer { //menos id tamaño total de campos identicos
			message = "err datos no coiciden con el modelo de la Tabla"
			return
		}
	case "all":
		if len(objectIn.FieldValue()) != len((*modelTableObject)) { //campos identicos
			message = "err datos no coiciden con el modelo de la Tabla"
			return
		}
	}
	wrongFields = verificationFieldToField(objectIn, *modelTableObject)

	if len(wrongFields) == 0 {
		ok = true
	} else {
		defMessage := `campo incorrecto`
		if len(wrongFields) > 1 {
			defMessage = `campos incorrectos`
		}

		wf := strings.Join(wrongFields, ", ")
		message = fmt.Sprintf("%v %v", wf, defMessage)
	}

	return
}

func verificationFieldToField(objectIn Object, modelTableObject map[string]model.Field) (wrongFields []string) {

	for fieldNameIn, dataIn := range objectIn.FieldValue() {

		if fieldModel, exists := (modelTableObject)[fieldNameIn]; exists {

			if fieldModel.Validate { // se requiere validar
				if !fieldModel.Input.Validate(dataIn) {
					wrongFields = append(wrongFields, fieldNameIn)
				}
			}

		} else {
			wrongFields = append(wrongFields, fieldNameIn)
		}
	}

	return
}
