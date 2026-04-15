package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	StepLength = 0.00075 // длина шага в км
)

// parsePackage парсит строку формата "1000,2h30m0s"
func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return steps, duration, nil
}

// DayActionInfo формирует итоговую строку
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return "" // Если ошибка, возвращаем пустую строку или текст ошибки
	}

	dist := float64(steps) * StepLength

	return fmt.Sprintf("Количество шагов: %d. Дистанция: %.2f км. Время в пути: %.2f ч",
		steps, dist, duration.Hours())
}
