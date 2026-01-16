package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			ErrInfoParse := fmt.Errorf("Error in dataset parsing: %w", err)
			log.Println(ErrInfoParse)
			continue
		}
		actionString, err := dp.ActionInfo()
		if err != nil {
			ErrInfo := fmt.Errorf("Error while outputting information: %w", err)
			log.Println(ErrInfo)
			continue
		}
		fmt.Println(actionString)
	}
}
