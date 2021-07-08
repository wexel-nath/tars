package db

import (
	"fmt"
	"strings"
)

type Updater struct{
	fieldsToUpdate map[string]interface{}
}

func NewUpdater() Updater {
	return Updater{
		fieldsToUpdate: map[string]interface{}{},
	}
}

func (u Updater) ShouldUpdate() bool {
	return len(u.fieldsToUpdate) > 0
}

func (u Updater) Set(field string, value interface{}) Updater {
	u.fieldsToUpdate[field] = value
	return u
}

func (u Updater) Output(placeholder int) ([]interface{}, string) {
	setParts := make([]string, 0)
	params := make([]interface{}, 0)
	for key, value := range u.fieldsToUpdate {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, placeholder))
		placeholder++
		params = append(params, value)
	}

	return params, strings.Join(setParts, ", ")
}
