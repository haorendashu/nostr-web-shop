package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadJsonConf(name string, itf interface{}) error {
	var e error
	_, err := os.Stat(name)
	if !os.IsNotExist(err) {
		bs, err := ioutil.ReadFile(name)
		if err != nil {
			err = errors.New(fmt.Sprintf("LoadJsonConf ioutil.ReadFile error %v", err.Error()))
		} else {
			err := json.Unmarshal(bs, itf)
			if err != nil {
				err = errors.New(fmt.Sprintf("json.Unmarshal error %v", err.Error()))
			}
		}
	} else {
		return err
	}

	return e
}

func LoadPlainArray(name string) (result []string, e error) {
	_, err := os.Stat(name)
	if !os.IsNotExist(err) {
		bs, err := ioutil.ReadFile(name)
		if err != nil {
			err = errors.New(fmt.Sprintf("LoadPlainArray error %v", err.Error()))
		} else {
			sc := bufio.NewScanner(bytes.NewReader(bs))
			for sc.Scan() {
				result = append(result, sc.Text())
			}
		}
	} else {
		e = err
	}

	return
}

func LoadPlain(name string) (str string, e error) {
	_, err := os.Stat(name)
	if !os.IsNotExist(err) {
		bs, err := ioutil.ReadFile(name)
		if err != nil {
			err = errors.New(fmt.Sprintf("LoadPlainArray error %v", err.Error()))
		} else {
			str = string(bs)
		}
	} else {
		e = err
	}

	return
}
