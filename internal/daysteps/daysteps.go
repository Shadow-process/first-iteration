package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid data format")
	}

	stepsStr := strings.TrimSpace(parts[0])
	durationStr := strings.TrimSpace(parts[1])

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps: %w", err)
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration: %w", err)
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) (string, error) {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return "", err
	}

	// длина шага зависит от роста (примерная формула)
	stepLength := height * 0.415 / 1000 // км

	distance := float64(steps) * stepLength

	result := fmt.Sprintf(
		"Количество шагов: %d. Дистанция: %.2f км. Время в пути: %.2f ч",
		steps,
		distance,
		duration.Hours(),
	)

	return result, nil
}
