package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/transform-in/xraw"
)

type dummy struct {
	ID        int
	Name      string
	Address   string
	IsActive  bool
	CreatedAt time.Time
}

type dummyOps struct {
	multi []*dummy
}

type Domain struct {
	GlobalRank     int
	TldRank        int
	Domain         string
	TLD            string
	RefSubNets     int64
	RefIPs         int64
	IDNDomain      string
	IDNTLD         string
	PrevGlobalRank int
	PrevTldRank    int
	PrevRefSubNets int64
	PrefRefIPs     int64
}

type readData struct {
	dummyCols []interface{}
	results   []map[string]interface{}
	single    map[string]interface{}
	model     interface{}
}

func main() {

	dbConfig := &xraw.DbConfig{
		Host:              "localhost",
		Username:          "root",
		Password:          "toor",
		DbName:            "customer_service",
		Driver:            "mysql",
		MaxIdleConnection: 10,
		MaxDBConnection:   10,
	}

	db, err := xraw.New(dbConfig)
	if err != nil {
		fmt.Errorf("Cannot connect to db")
		os.Exit(1)
	}
	fmt.Println("success to connect", db)

}

func generateDummyColumn() <-chan readData {
	results := make(chan readData)
	colLength := 5
	go func() {
		var res []map[string]interface{}
		results <- readData{
			dummyCols: make([]interface{}, colLength),
			results:   res,
		}
		close(results)
	}()

	return results
}

func readEachColumnAndStore(results []dummy, dataIn <-chan readData) <-chan readData {
	wg := new(sync.WaitGroup)
	wg.Add(len(results))

	resultChan := make(chan []map[string]interface{})
	var mtx sync.Mutex

	go func() {
		var multi []map[string]interface{}
		for _, result := range results {
			fmt.Println("read: ", result)
			//go func(result dummy, multi []map[string]interface{}, mtx sync.Mutex) {
			mtx.Lock()
			bByte, _ := json.Marshal(result)
			single := make(map[string]interface{})
			//go func(byteResult []byte, mapDest map[string]interface{}) {
			json.Unmarshal(bByte, &single)
			multi = append(multi, single)
			mtx.Unlock()
			wg.Done()
			//}(result, multi, mtx)
		}
		fmt.Println("store the result")
		resultChan <- multi
	}()

	go func() {
		wg.Wait()
		for data := range resultChan {
			fmt.Println(data)
		}
		close(resultChan)
	}()

	return make(chan readData)
}

func readEacColumn(results []dummy, dataIn <-chan readData) <-chan readData {
	chanResult := make(chan readData)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var multiData []map[string]interface{}
		for data := range dataIn {
			multiData = data.results
			fmt.Println("total data", len(results))
			wg.Add(len(results))
			go func(multiData []map[string]interface{}) {
				for _, result := range results {
					singleData := make(map[string]interface{})
					bBytes, _ := json.Marshal(result)

					json.Unmarshal(bBytes, &singleData)
					multiData = append(multiData, singleData)
					fmt.Println("result", result, singleData, len(multiData))
					wg.Done()
				}
				fmt.Println("inseide:", len(multiData))

			}(multiData)
			fmt.Println(len(multiData))
		}
		chanResult <- readData{
			results: multiData,
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		fmt.Println("finished executed")
		close(chanResult)

	}()

	return chanResult
}

func readEachColumnResult(results []*dummy, dataIn <-chan readData) <-chan readData {
	chanResult := make(chan readData)
	wg := new(sync.WaitGroup)

	//wg.Add(len(results))
	//multiResult := make([]map[string]interface{}, len(results))

	//mtx := new(sync.Mutex)
	go func() {
		for _, result := range results {
			wg.Add(1)
			bRes, _ := json.Marshal(result)
			single := make(map[string]interface{})

			go func(byteResult []byte, destMap map[string]interface{}) {
				for data := range dataIn {
					json.Unmarshal(byteResult, &single)
					data.results = append(data.results, single)

					chanResult <- data
				}
				wg.Done()
			}(bRes, single)

		}
	}()
	//fmt.Println(multiResult)
	wg.Wait()
	close(chanResult)
	//}()
	return chanResult
}

func (d *dummyOps) generateDummyMultiple() []*dummy {
	var dummyMultipleRows []*dummy
	for x := 0; x < 1000; x++ {
		dummyMultipleRows = append(dummyMultipleRows, &dummy{
			ID:        x + 1,
			Name:      "test",
			Address:   "test",
			IsActive:  false,
			CreatedAt: time.Now(),
		})
	}

	d.multi = dummyMultipleRows
	return dummyMultipleRows
}

func getNoOfWorker(totalData int) int {
	switch {
	case totalData > 2 && totalData <= 5:
		return 2
	case totalData <= 10:
		return totalData / 2
	case totalData <= 20:
		return 5
	case totalData <= 100:
		return 10
	case totalData > 100:
		return 20
	default:
		return 1
	}
}
