package repository

import (
	"time"

	"github.com/google/wire"
)

var _ TimeRepository = (*timeRepository)(nil)

var TimeRepositorySet = wire.NewSet(
	NewTimeRepositoy,
	wire.Bind(new(TimeRepository), new(*timeRepository)),
)

type TimeRepository interface {
	Now() time.Time
}

type timeRepository struct{}

func NewTimeRepositoy() TimeRepository {
	return &timeRepository{}
}

var _ TimeRepository = (*timeRepository)(nil)

func (t *timeRepository) Now() time.Time {
	return time.Now()
}
