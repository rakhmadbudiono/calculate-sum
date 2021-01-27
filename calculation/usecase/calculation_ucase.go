package usecase

import (
	"context"
	"time"

	"github.com/rakhmadbudiono/calculate-sum/domain"
)

type calculationUsecase struct {
	contextTimeout time.Duration
}

// NewCalculationUsecase will create new an calculationUsecase object representation of domain.NewCalculationUsecase interface
func NewCalculationUsecase(timeout time.Duration) domain.CalculationUsecase {
	return &calculationUsecase{
		contextTimeout: timeout,
	}
}

/*
* In this function below, I'm using errgroup with the pipeline pattern
* Look how this works in this package explanation
* in godoc: https://godoc.org/golang.org/x/sync/errgroup#ex-Group--Pipeline
 */
func (s *calculationUsecase) Add(c context.Context, a int64, b int64) (res int64, err error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	res = a + b
	return
}
