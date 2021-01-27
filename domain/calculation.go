package domain

import "context"

// Calculation ...
type Calculation struct {}

// CalculationUsecase represent the calculation's usecases
type CalculationUsecase interface {
	Add(ctx context.Context, a int64, b int64) (int64, error)
}
