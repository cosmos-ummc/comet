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
	Job       func(ctx context.Context, d string) error
	RevokeJob func(ctx context.Context) error
	Wg        sync.WaitGroup
}

func (it *Scheduler) isr() {
	fmt.Printf("Scheduler : Job called\n")
	if it.Enabled {
		now := MalaysiaTime(time.Now())
		if it.HaveStart {
			ctx := context.Background()

			// generate report
			err := it.Job(ctx, TimeToDateString(now))
			if err != nil {
				logger.Log.Error("error in scheduler job function", zap.String("reason", err.Error()))
			}

			// force revoke all user tokens
			err = it.RevokeJob(ctx)
			if err != nil {
				logger.Log.Error("error in scheduler revoke job function", zap.String("reason", err.Error()))
			}

			ctx.Done()
		} else {
			it.HaveStart = true
		}
		t, _ := DateStringToTime(TimeToDateString(now.Add(24 * time.Hour)))
		// change to local time 0001
		t = t.Add(1 * time.Minute)
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
