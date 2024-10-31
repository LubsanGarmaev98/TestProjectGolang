package service

import "time"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	return int64(time.Until(date).Hours()) / 24
}
