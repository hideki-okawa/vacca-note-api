package main

import (
	"fmt"
	"net/http"

	"github.com/Okaki030/vacca-note-server/log"
	"github.com/najeira/measure"
)

// type StatsLog struct {
// 	Key   string
// 	Count int64
// 	Sum   float64
// 	Min   float64
// 	Max   float64
// 	Avg   float64
// 	Rate  float64
// 	P95   float64
// }

// getStats はAPIの実行時間を出力します。
func getStats(w http.ResponseWriter, req *http.Request) {
	log.Debugf("run getStats")

	stats := measure.GetStats()
	stats.SortDesc("sum")

	// var statsLogs []StatsLog
	fmt.Printf("key, count, sum, avg\n")
	for _, s := range stats {
		// statsLog := StatsLog{
		// 	Key:   s.Key,
		// 	Count: s.Count,
		// 	Sum:   s.Sum,
		// 	Min:   s.Min,
		// 	Max:   s.Max,
		// 	Avg:   s.Avg,
		// 	Rate:  s.Rate,
		// 	P95:   s.P95,
		// }
		// statsLogs = append(statsLogs, statsLog)
		fmt.Printf("%s, %d, %f, %f \n", s.Key, s.Count, s.Sum, s.Avg)
		// fmt.Printf("%s,%d,%f,%f,%f,%f,%f,%f\n",
		// 	s.Key, s.Count, s.Sum, s.Min, s.Max, s.Avg, s.Rate, s.P95)
	}
}

func deleteStats(w http.ResponseWriter, req *http.Request) {
	log.Debugf("run deleteStats")

	measure.Reset()
	log.Debugf("complete reset measure")
}
