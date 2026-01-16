package trainings

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Custom errors
var (
	ErrInvalidFormat          = fmt.Errorf("invalid data format")
	ErrStepsMustBePositive    = fmt.Errorf("the number of steps must be positive")
	ErrDurationMustBePositive = fmt.Errorf("duration must be positive")
	ErrInvalidActivityType    = fmt.Errorf("неизвестный тип тренировки")
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")
	if len(dataSlice) != 3 {
		log.Println(ErrInvalidFormat)
		return fmt.Errorf("Training parsing failed: %w", ErrInvalidFormat)
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
	parsedActivity := dataSlice[1]

	parsedDuration, err := time.ParseDuration(dataSlice[2])
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Duration parsing failed: %w", err)
	}
	if parsedDuration <= 0 {
		log.Println(ErrDurationMustBePositive)
		return fmt.Errorf("Invalid duration: %w", ErrDurationMustBePositive)
	}

	t.Steps = parsedSteps
	t.TrainingType = parsedActivity
	t.Duration = parsedDuration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var err error
	var calories float64
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	if t.TrainingType == "Бег" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			ErrRunningCalories := fmt.Errorf("Invalid calories data for ActionInfo: %w", err)
			log.Println(ErrRunningCalories)
			return "", ErrRunningCalories
		}
	} else if t.TrainingType == "Ходьба" {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			ErrWalkingCalories := fmt.Errorf("Invalid calories data for ActionInfo: %w", err)
			log.Println(ErrWalkingCalories)
			return "", ErrWalkingCalories
		}
	} else {
		log.Println(ErrInvalidActivityType)
		return "", fmt.Errorf("Invalid activity data for ActionInfo: %w", ErrInvalidActivityType)
	}
	dataString := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)
	return dataString, nil

}
