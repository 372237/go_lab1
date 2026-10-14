package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps содержит данные о дневной активности (прогулке).
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse парсит строку с данными о прогулке и заполняет структуру.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid data format: expected 2 fields, got %d", len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("failed to parse steps: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("steps must be positive")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("failed to parse duration: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("duration must be positive")
	}
	ds.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	info += fmt.Sprintf("Дистанция составила %.2f км.\n", distance)
	info += fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories)

	return info, nil
}
