package xraw

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"sync"
)

type selectWorker struct {
	rows   *sqlx.Rows
	cols   []string
	data   chan []interface{}
	wg     *sync.WaitGroup
	result interface{}
}

func (xe *Engine) scanData(dt *selectWorker) {
	dt.data = make(chan []interface{}, 0)
	for dt.rows.Next() {
		var tmpData []interface{}
		dt.rows.Scan(tmpData...)

		dt.data <- tmpData
	}
	close(dt.data)
}

func (xe *Engine) executeWorker(dt *selectWorker) {
	multipleResult := make(chan []map[string]interface{})
	var multiResult []map[string]interface{}

	go func(multiResData <-chan []map[string]interface{}, model interface{}) {
		for data := range multiResData {
			var willBeMarshal interface{}
			switch {
			case len(data) == 1:
				willBeMarshal = data[0]
			case len(data) == 0:
				model = nil
				return
			default:
				willBeMarshal = data
			}

			bJSON, _ := json.Marshal(willBeMarshal)
			json.Unmarshal(bJSON, model)
		}
	}(multipleResult, dt.result)

	for workerIndex := 0; workerIndex < xe.workers.total; workerIndex++ {
		go func(wi int, cols []string, dt <-chan []interface{}, wg *sync.WaitGroup) {
			for filledData := range dt {
				// job run
				singleResult := make(map[string]interface{})
				for idx, value := range filledData {
					singleResult[cols[idx]] = value
				}
				multiResult = append(multiResult, singleResult)
				wg.Done()
			}
		}(workerIndex, dt.cols, dt.data, dt.wg)
	}
	multipleResult <- multiResult
	close(multipleResult)
}
