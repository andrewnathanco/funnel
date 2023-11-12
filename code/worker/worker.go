package worker

import (
	"fmt"
	"funnel/db"
	"log/slog"
)

// need to do everything as utc
func (s FunnelSchedular) RegisterWorkers() error {
	// register session worker
	// get current page
	funnel_meta, err := db.FunnelDAL.GetFunnelMeta()
	if err != nil {
		slog.Error(fmt.Errorf("could not get current movie page: %w", err).Error())
		return err
	}

	// tmdb can't go past 500 so we don't need to cache anymore
	if funnel_meta.CurrentTMDBPage < 500 {
		s.Scheduler.Every(1).Minute().Do(s.TMDBWorker.CacheAnswers)
	}

	return nil
}

func (s FunnelSchedular) InitProgram() {
	// need to manually pull a few answers to start
	funnel_meta, err := db.FunnelDAL.GetFunnelMeta()
	if err != nil {
		slog.Error(fmt.Errorf("could not get current movie page: %w", err).Error())
	}

	// tmdb can't go past 500 so we don't need to cache anymore
	if funnel_meta.CurrentTMDBPage < 500 {
		s.TMDBWorker.CacheAnswers()
	}
}
