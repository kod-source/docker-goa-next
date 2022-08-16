package repository

import "time"

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
