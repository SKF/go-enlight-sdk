package models

import "github.com/SKF/go-utility/uuid"

func stringToUUIDArray(stringArr []string) (uuidArr []uuid.UUID) {
	for _, elem := range stringArr {
		uuidArr = append(uuidArr, uuid.UUID(elem))
	}
	return
}

func uuidToStringArray(uuidArr []uuid.UUID) (stringArr []string) {
	for _, elem := range uuidArr {
		stringArr = append(stringArr, elem.String())
	}
	return
}
