package migrate

import (
	"fmt"
	"strings"

	"github.com/cdvelop/margo/storage/share"
)

func createSqlListByField(tableName string, fields ...map[string]string) (sqlist, keyList []string) {
	var (
		defaulType     string   //TEXT tipo por defecto en base de datos
		foreignKeyList []string //FOREIGN KEY si los hay
	)

	for _, field := range fields {

		for keyName, valueType := range field {
			defaulType = ` TEXT`
			if valueType != "" {
				defaulType = ` ` + valueType
			}

			if primaryKey, primaryKeyThisTable := share.IdpkTABLA(keyName, tableName); primaryKey {
				if primaryKeyThisTable {
					defaulType = defaulType + ` PRIMARY KEY NOT NULL`
				}
				if !primaryKeyThisTable { //id pero son FOREIGN KEY
					defaulType = defaulType + ` NOT NULL`
					fkName := keyName[2:]                                   //nombre tabla foranea
					foreignTableName := strings.ReplaceAll(fkName, "_", "") //remover _
					cf := fmt.Sprintf("CONSTRAINT fk_%v FOREIGN KEY (%v) REFERENCES %v(%v)", foreignTableName, keyName, foreignTableName, keyName)
					// CONSTRAINT fk_departments FOREIGN KEY (department_id) REFERENCES departments(department_id)
					foreignKeyList = append(foreignKeyList, cf)
				}
			}

			sqlist = append(sqlist, keyName+defaulType)
			keyList = append(keyList, keyName)
		}
	}

	//hay FOREIGN KEY ?
	if len(foreignKeyList) > 0 {
		c := `` //coma si es necesario
		if len(foreignKeyList) > 1 {
			c = `, `
		}
		f := strings.Join(foreignKeyList, c)
		sqlist = append(sqlist, f)
	}
	return
}
