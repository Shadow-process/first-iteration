package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	WalkingCaloriesWeightMultiplier = 0.035
	WalkingSpeedHeightMultiplier    = 0.029
)

// SpentCalories рассчитывает калории для ходьбы
func SpentCalories(data string, weight, height float64) (float64, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, err
	}

	// Средняя скорость в км/ч
	durationHours := duration.Hours()
	if durationHours == 0 {
		return 0, nil
	}
	meanSpeed := (float64(steps) * 0.00075) / durationHours

	// Формула из задания
	calories := (WalkingCaloriesWeightMultiplier*weight + (meanSpeed*meanSpeed/height)*WalkingSpeedHeightMultiplier*weight) * (durationHours * 60)

	return calories, nil
}
