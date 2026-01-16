package spentenergy

import (
	"fmt"
	"log"
	"time"
)

// Custom errors
var (
	DataBelowZero = fmt.Errorf("data below zero")
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		log.Println(DataBelowZero)
		return 0, fmt.Errorf("Invalid data for WalkingSpentCalories: %w", DataBelowZero)
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	caloriesWalking := ((weight * meanSpeed * duration.Minutes()) / minInH) * walkingCaloriesCoefficient
	return caloriesWalking, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		log.Println(DataBelowZero)
		return 0, fmt.Errorf("Invalid data for RunningSpentCalories: %w", DataBelowZero)
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	calories := (weight * meanSpeed * duration.Minutes()) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		log.Println(DataBelowZero)
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		log.Println(DataBelowZero)
		return 0
	}
	stepLength := height * stepLengthCoefficient
	distance := (float64(steps) * stepLength) / mInKm
	return distance
}
