package util

import (
	goTime "time"
)

type Time interface {
	GetCurrent() goTime.Time
}

type time struct {
}

func (util *time) GetCurrent() goTime.Time {
	return goTime.Now()
}

func NewTime() Time {
	return &time{}
}
