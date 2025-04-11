package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{})

	for _, jobb := range jobs {

		out := make(chan interface{})
		wg.Add(1)

		go func(job job, in, out chan interface{}) {
			defer wg.Done()
			defer close(out)
			job(in, out)
		}(jobb, in, out)

		in = out
	}

	wg.Wait()
}

func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	chanMd5in := make(chan string)
	chanMd5out := make(chan string)

	go func(chanin chan string, chanout chan string) {
		for str := range chanin {
			chanout <- DataSignerMd5(str)
		}
	}(chanMd5in, chanMd5out)

	for rawdata := range in {
		wg.Add(1)
		data := fmt.Sprintf("%v", rawdata)

		datamd5 := DataSignerMd5(data)

		go func() {
			defer wg.Done()

			c32 := make(chan string)
			md5 := make(chan string)

			go func(data32 string) {
				c32 <- DataSignerCrc32(data32)
			}(data)

			go func(data5 string) {
				md5 <- DataSignerCrc32(data5)
			}(datamd5)

			datafrom32 := <-c32
			datafrom5 := <-md5

			out <- datafrom32 + "~" + datafrom5

		}()
	}
	wg.Wait()
	close(chanMd5in)

}

func MultiHash(in, out chan interface{}) {

	wg := &sync.WaitGroup{}

	for datainterface := range in {
		data := datainterface.(string)

		wg.Add(1)

		go func() {
			defer wg.Done()

			var (
				//builder  strings.Builder
				results  = make([]string, 6)
				wgHashes sync.WaitGroup
				//mu       sync.Mutex
			)

			wgHashes.Add(6)
			for th := 0; th < 6; th++ {
				go func(t int) {
					defer wgHashes.Add(-1)
					hash := DataSignerCrc32(strconv.Itoa(t) + data)
					//mu.Lock()
					results[t] = hash
					//mu.Unlock()
				}(th)
			}

			wgHashes.Wait()

			out <- strings.Join(results, "")

		}()
	}

	wg.Wait()

}
func CombineResults(in, out chan interface{}) {

	result := []string{}
	for data := range in {
		result = append(result, data.(string))

	}
	sort.Strings(result)
	out <- strings.Join(result, "_")
}
