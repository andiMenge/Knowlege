package main

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"strings"
)

func main() {

	//new WfsJobs instance
	wfsJobs := new(WfsJobs)

	//get Job Guids
	err := wfsJobs.getGuids()
	checkErr(err)

	//get details from Job(GUID) and add Job to WfsJobs.Jobs
	err = wfsJobs.getJobDetails()
	checkErr(err)

	//print count of jobs in WfsJobs.Jobs Array
	fmt.Println("Jobs:", wfsJobs.countJobs())

	//format to json and print to stdout
	json, err := wfsJobs.toJson()
	checkErr(err)
	fmt.Println(string(json))
}

type WfsJob struct {
	Name   string
	Guid   string
	ErrMsg string
}

type WfsJobs struct {
	Guids []string `json:"-"`
	Jobs  []WfsJob `json:"wfsJobs"`
}

func (w *WfsJobs) getGuids() error {
	uuids := uuidGen(10) //mock of guid src
	if len(uuids) == 0 {
		return fmt.Errorf("getGuids(): no guids in array (len == 0)")
	}
	//add Guids to WfsJobs.Guids
	for _, i := range uuids {
		w.Guids = append(w.Guids, i)
	}
	return nil
}

func (w *WfsJobs) getJobDetails() error {
	//read json file
	path := "/Users/andi/Documents/tmp/wfsJob.json"
	data, err := readFile(path)
	if err != nil {
		return fmt.Errorf("readFile failed: ", err)
	}

	for _, i := range w.Guids {
		//make http call to wfs api to get json
		//new WfsJob instance
		job := new(WfsJob)

		//parse json to WfsJob instance
		err := json.Unmarshal(data, &job)
		if err != nil {
			return fmt.Errorf("error parsing json", err)
		}

		job.ErrMsg = getWfsJobErrMsg(data) //set ErrMsg attribute

		w.Jobs = append(w.Jobs, *job)

		//temp workaraound
		_ = i
		println("added Job")
	}
	return nil
}

func (w *WfsJobs) countJobs() int {
	return len(w.Jobs)
}

func (w *WfsJobs) toJson() ([]byte, error) {
	json, err := json.Marshal(w)
	if err != nil {
		return nil, fmt.Errorf("Error while encoding to json:", err)
	}
	return json, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error while reading file: ", err)
	}
	return data, nil
}

//takes json as byte array as input; returns msg as string
func getWfsJobErrMsg(data []byte) string {
	wfsErrMsg := ""
	//Get the Task Property with Name=Error from the wfs json response
	searchResult := gjson.GetBytes(data, `Task.#.Property.#[Name="Error"]`)
	taskPropErrStr := searchResult.String()

	//replace "[" and "]" from string
	cleanTaskPropErr := strings.Replace(taskPropErrStr, "[", "", -1)
	cleanTaskPropErr = strings.Replace(cleanTaskPropErr, "]", "", -1)

	//Search for "Value" in the "cleanTaskPropErr" json object
	if gjson.Get(cleanTaskPropErr, "Value").Exists() { //Check if Err Msg in json exists
		m := gjson.Get(cleanTaskPropErr, "Value")

		if m.String() == "" {
			wfsErrMsg = "no errors"
		} else {
			wfsErrMsg = m.String()
		}
	}

	return wfsErrMsg
}

func uuidGen(count int) []string {
	uuids := make([]string, 0)
	for i := 0; i < count; i++ {
		u := uuid.NewV4()
		uuids = append(uuids, string(u[:]))
	}
	return uuids
}
