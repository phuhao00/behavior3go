package actions

import (
	b3 "github.com/phuhao00/behavior3go"
	. "github.com/phuhao00/behavior3go/core"
)

type Error struct {
	Action
}

func (this *Error) OnTick(tick *Tick) b3.Status {
	return b3.ERROR
}
