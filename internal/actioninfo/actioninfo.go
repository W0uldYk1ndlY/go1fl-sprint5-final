package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println(err)
		}
		result, err := dp.ActionInfo()
		fmt.Println(result)
		if err != nil {
			log.Println(err)
		}

	}
}
