package utility

import (
	"comet/pkg/logger"
	"context"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Scheduler struct {
	Enabled   bool
	HaveStart bool
	Job       func(ctx context.Context) error
	Wg        sync.WaitGroup
}

func (it *Scheduler) isr() {
	fmt.Printf("Scheduler : Job called\n")
	if it.Enabled {
		now := MalaysiaTime(time.Now())
		if it.HaveStart {
			ctx := context.Background()

			// generate report
			err := it.Job(ctx)
			if err != nil {
				logger.Log.Error("error in scheduler job function", zap.String("reason", err.Error()))
			}

			ctx.Done()
		} else {
			it.HaveStart = true
		}
		t, _ := DateStringToTime(TimeToDateString(now))
		// change to local time 0900 -> 1000
		t = t.Add(10 * time.Hour)
		fmt.Printf("\tJob interval %v\n", t.Sub(now))
		time.AfterFunc(t.Sub(now), it.isr)
	} else {
		it.Wg.Done()
	}
}

//trigger
func (it *Scheduler) Start() {
	if it.Enabled {
		it.Wg.Add(1)
		//now := utility.MalaysiaTime(time.Now())
		//t, _ := DateStringToTime(TimeToDateString(now))
		//time.AfterFunc(t.Sub(now), it.isr)
		it.isr()
	}
}
