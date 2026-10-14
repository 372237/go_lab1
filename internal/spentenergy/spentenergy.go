package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Distance рассчитывает дистанцию в километрах по количеству шагов и росту.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

// MeanSpeed рассчитывает среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

// RunningSpentCalories рассчитывает калории, потраченные при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be positive")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be positive")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be positive")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be positive")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

// WalkingSpentCalories рассчитывает калории, потраченные при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be positive")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be positive")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be positive")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be positive")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories * walkingCaloriesCoefficient, nil
}
