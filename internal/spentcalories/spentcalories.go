package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const stepLength = 0.00075 // км

func SpentCalories(data string, weight, height float64) (float64, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid data format")
	}

	stepsStr := strings.TrimSpace(parts[0])
	durationStr := strings.TrimSpace(parts[1])

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, fmt.Errorf("invalid steps: %w", err)
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, fmt.Errorf("invalid duration: %w", err)
	}

	durationHours := duration.Hours()
	if durationHours <= 0 {
		return 0, fmt.Errorf("duration must be greater than zero")
	}

	meanSpeed := (float64(steps) * stepLength) / durationHours

	calories := (0.035*weight + (meanSpeed*meanSpeed/height)*0.029*weight) * durationHours * 60

	return calories, nil
}
