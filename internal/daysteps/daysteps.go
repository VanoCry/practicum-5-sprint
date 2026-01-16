package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var (
	ErrInvalidFormat          = fmt.Errorf("invalid data format")
	ErrStepsMustBePositive    = fmt.Errorf("the number of steps must be positive")
	ErrDurationMustBePositive = fmt.Errorf("duration must be positive")
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")
	if len(dataSlice) != 2 {
		log.Println(ErrInvalidFormat)
		return fmt.Errorf("DaySteps parsing failed: %w", ErrInvalidFormat)
	}
	parsedSteps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Steps parsing failed: %w", err)
	}
	if parsedSteps <= 0 {
		log.Println(ErrStepsMustBePositive)
		return fmt.Errorf("Invalid steps: %w", ErrStepsMustBePositive)
	}

	parsedDuration, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Duration parsing failed: %w", err)
	}
	if parsedDuration <= 0 {
		log.Println(ErrDurationMustBePositive)
		return fmt.Errorf("Invalid duration: %w", ErrDurationMustBePositive)
	}
	ds.Steps = parsedSteps
	ds.Duration = parsedDuration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	/*Количество шагов: 792.
	Дистанция составила 0.51 км.
	Вы сожгли 221.33 ккал.*/
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("Error in calorie calculation: %w", err)
	}
	dataString := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories)
	return dataString, nil
}
