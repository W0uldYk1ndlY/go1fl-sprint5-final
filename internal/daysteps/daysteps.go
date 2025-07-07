package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parse := strings.Split(datastring, ",")

	if len(parse) != 2 {
		return fmt.Errorf("daysteps len error: len != 2")
	}

	steps, err := strconv.Atoi(parse[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("error: steps <= 0")
	}

	ds.Steps = steps

	duration, err := time.ParseDuration(parse[1])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}

	if duration <= 0 {
		return fmt.Errorf("errod: duration <= 0")
	}

	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
