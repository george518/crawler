/************************************************************
** @Description: model
** @Author: haodaquan
** @Date:   2018-03-17 22:22
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-17 22:22
*************************************************************/
package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hukou      string
	Xingzuo    string
	House      string
	Car        string
	City       string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
