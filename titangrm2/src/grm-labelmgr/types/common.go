package types

import ()

type UserData struct {
	Name     string `json:"name"`
	Data     string `json:"data"`
	DataType string `json:"data_type"`
	UserId   string `json:"user_id,omitempty"`
}

type UserDataList []UserData
