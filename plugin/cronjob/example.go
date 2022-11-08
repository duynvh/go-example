package cronjob

import (
	"github.com/200Lab-Education/go-sdk/logger"
	"time"
)

type myCronJob struct {
	logger logger.Logger
}

func NewMyCronJob() *myCronJob {
	return &myCronJob{}
}

func (myCronJob) GetPrefix() string {
	return "my-cron-job"
}

func (j *myCronJob) Get() interface{} {
	return j
}

func (myCronJob) Name() string {
	return "my-cron-job"
}

func (myCronJob) InitFlags() {

}

func (j *myCronJob) Configure() error {
	j.logger = logger.GetCurrent().GetLogger("cronjob")
	return nil
}

func (j *myCronJob) Run() error {
	if err := j.Configure(); err != nil {
		return err
	}

	for i := 1; i <= 100; i++ {
		time.Sleep(time.Second)
		j.logger.Infoln(i)
	}

	return nil
}

func (myCronJob) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}
