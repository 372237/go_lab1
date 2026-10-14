package actioninfo

import "log"

// DataParser описывает методы, которые должны быть реализованы
// структурами Training и DaySteps.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info обрабатывает слайс строк с данными и выводит информацию
// об активностях, используя переданный DataParser.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("ошибка парсинга данных: %v", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования информации: %v", err)
			continue
		}

		log.Println(info)
	}
}
