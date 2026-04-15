func SpentCalories(data string, weight, height float64) (float64, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		// Убери точку в конце сообщения "invalid data format"
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

	durationHours := duration.Hours()
	if durationHours == 0 {
		return 0, nil
	}

	meanSpeed := (float64(steps) * 0.00075) / durationHours

	// Вот ПРАВИЛЬНАЯ формула (добавь * weight во вторую часть):
	calories := (0.035*weight + (meanSpeed*meanSpeed/height)*0.029*weight) * durationHours * 60

	return calories, nil
}