# Golang Docker
Compile static linked binary

```
export CGO_ENABLED=0
export GOOS=linux 
go build -a -installsuffix cgo -o binaryName .
```

## Docker Stuff

Build Container

```docker build -t compiledBinary -f Dockerfile .```


# Get fields from complex json without parsing

github.com/tidwall/gjson

**replace "" in line 11 with ``**

```
//takes json as byte array as input; returns msg as string
func getWfsJobErrMsg(data []byte) string {
	//Get the Task Property with Name=Error from the wfs json response
	searchResult := gjson.GetBytes(data, "Task.#.Property.#[Name="Error"]")
	taskPropErrStr := searchResult.String()

	//replace "[" and "]" from string
	cleanTaskPropErr := strings.Replace(taskPropErrStr, "[", "", -1)
	cleanTaskPropErr = strings.Replace(cleanTaskPropErr, "]", "", -1)

	//Search for "Value" in the "cleanTaskPropErr" json object
	errMsg := gjson.Get(cleanTaskPropErr, "Value")

	return errMsg.String()
}
```

# Generate UUIDs

``` "github.com/satori/go.uuid" ```

```
func uuidGen(count int) []string {
	uuids := make([]string, 0)
	for i := 0; i < count; i++ {
		u := uuid.NewV4()
		uuids = append(uuids, string(u[:])) //convert UUID type to string and append to slice
	}
	return uuids
}
```

# Parse WFS Job example

```
package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"strings"
)

type WfsJobs struct {
	Jobs []WfsJob `json:"wfsJobs"`
}

func (w *WfsJobs) addJob(job *WfsJob) {
	w.Jobs = append(w.Jobs, *job)
}

type WfsJob struct {
	Name   string
	Guid   string
	ErrMsg string
}

func main() {
	//read json file
	path := "/Users/andi/Documents/tmp/wfsJob.json"
	data := readFile(path)

	//new WfsJob instance
	wfsJob := new(WfsJob)

	//parse Name and Guid from json to wfsJob instance
	json.Unmarshal(data, &wfsJob)

	//set ErrMsg attribute
	wfsJob.ErrMsg = getWfsJobErrMsg(data)

	//new WfsJobs instance
	wfsJobs := new(WfsJobs)
	wfsJobs.addJob(wfsJob) //add wfsJob to []Jobs

	//fmt.Println(wfsJobs)

	json, err := json.Marshal(wfsJobs)
	checkErr(err)

	fmt.Println(string(json))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	checkErr(err)
	return data
}

//takes json as byte array as input; returns msg as string
func getWfsJobErrMsg(data []byte) string {
	//Get the Task Property with Name=Error from the wfs json response
	searchResult := gjson.GetBytes(data, `Task.#.Property.#[Name="Error"]`)
	taskPropErrStr := searchResult.String()

	//replace "[" and "]" from string
	cleanTaskPropErr := strings.Replace(taskPropErrStr, "[", "", -1)
	cleanTaskPropErr = strings.Replace(cleanTaskPropErr, "]", "", -1)

	//Search for "Value" in the "cleanTaskPropErr" json object
	errMsg := gjson.Get(cleanTaskPropErr, "Value")

	return errMsg.String()
}
```

# RestAPI 'http-Post' example

```
func Queuejob(w http.ResponseWriter, r *http.Request) {
	var job Job
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &job); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Fprintln(w, "name: ", job.Name)
}
```
# Rest Api Stuff

## RestApi return http error
```
if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
```

## Return 200 OK
```
w.WriteHeader(http.StatusOK)
```

