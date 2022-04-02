package data_loader

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type JsonFile struct {
	Json []Json
}

type Json struct {
	Date int64
	Num  int
}

func (j *JsonFile) Load() error {
	file , err := ioutil.ReadFile("/data_validation/data.json")
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(file, &j)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

