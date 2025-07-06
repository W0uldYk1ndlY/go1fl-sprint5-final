package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parse := strings.Split(datastring, ",")

	if len(parse) != 3 {
		return fmt.Errorf("Ошибка парсинга. Длина массива не 3")
	}

	steps, err := strconv.Atoi(parse[0])
	if err != nil {
		return fmt.Errorf("Ошибка парсинга. Шаги в число")
	}

	if steps <= 0 {
		return fmt.Errorf("Некорректное количество шагов")
	}

	t.Steps = steps
	t.TrainingType = parse[1]

	duration, err := time.ParseDuration(parse[2])
	if err != nil {
		return fmt.Errorf("Ошибка парсинга. Время в time.Duration")
	}

	if duration <= 0 {
		return fmt.Errorf("Некорректная продолжительность")
	}

	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.TrainingType != "Ходьба" && t.TrainingType != "Бег" {
		return "", fmt.Errorf("Некорректный тип тренировки")
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	if t.TrainingType == "Бег" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	} else {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil

}
