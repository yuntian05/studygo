package model

import "encoding/json"

type Profile struct {
	Name       string
	Marriage   string
	Age        string
	Xingzuo    string
	Height     string
	Weight     string
	Workplace     string
	Income     string
	Occupation string
	Education  string
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