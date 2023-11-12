package worker

import (
	"time"

	"github.com/go-co-op/gocron"
)

type FunnelSchedular struct {
	Scheduler  *gocron.Scheduler
	TMDBWorker TMDBWorker
}

func NewFunnelSchedular() *FunnelSchedular {
	tmdb_worker := TMDBWorker{}
	return &FunnelSchedular{
		Scheduler:  gocron.NewScheduler(time.UTC),
		TMDBWorker: tmdb_worker,
	}

}

func (fs FunnelSchedular) StartScheduler() {
	fs.Scheduler.StartAsync()
}
