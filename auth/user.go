// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"encoding/json"
	"fmt"
)

// AuthConfig is the core configuration.
// The first step to use this SDK is to use the InitConfig function to initialize the global authConfig.
type AuthConfig struct {
	Endpoint         string
	ClientId         string
	ClientSecret     string
	JwtSecret        string
	OrganizationName string
}

var authConfig AuthConfig

// User has the same definition as https://github.com/casbin/casdoor/blob/master/object/user.go#L24,
// used to obtain user-related information from your Casdoor server.
type User struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	UpdatedTime string `xorm:"varchar(100)" json:"updatedTime"`
	
	Id                string   `xorm:"varchar(100)" json:"id"`
	Type              string   `xorm:"varchar(100)" json:"type"`
	Password          string   `xorm:"varchar(100)" json:"password"`
	DisplayName       string   `xorm:"varchar(100)" json:"displayName"`
	Avatar            string   `xorm:"varchar(255)" json:"avatar"`
	Email             string   `xorm:"varchar(100)" json:"email"`
	Phone             string   `xorm:"varchar(100)" json:"phone"`
	Address           []string `json:"address"`
	Affiliation       string   `xorm:"varchar(100)" json:"affiliation"`
	Tag               string   `xorm:"varchar(100)" json:"tag"`
	Language          string   `xorm:"varchar(100)" json:"language"`
	Score             int      `json:"score"`
	IsAdmin           bool     `json:"isAdmin"`
	IsGlobalAdmin     bool     `json:"isGlobalAdmin"`
	IsForbidden       bool     `json:"isForbidden"`
	SignupApplication string   `xorm:"varchar(100)" json:"signupApplication"`
	Hash              string   `xorm:"varchar(100)" json:"hash"`
	PreHash           string   `xorm:"varchar(100)" json:"preHash"`
	
	Github   string `xorm:"varchar(100)" json:"github"`
	Google   string `xorm:"varchar(100)" json:"google"`
	QQ       string `xorm:"qq varchar(100)" json:"qq"`
	WeChat   string `xorm:"wechat varchar(100)" json:"wechat"`
	Facebook string `xorm:"facebook varchar(100)" json:"facebook"`
	DingTalk string `xorm:"dingtalk varchar(100)" json:"dingtalk"`
	Weibo    string `xorm:"weibo varchar(100)" json:"weibo"`
	Gitee    string `xorm:"gitee varchar(100)" json:"gitee"`
	LinkedIn string `xorm:"linkedin varchar(100)" json:"linkedin"`
	
	Properties map[string]string `json:"properties"`
}

func InitConfig(endpoint string, clientId string, clientSecret string, jwtSecret string, organizationName string) {
	authConfig = AuthConfig{
		Endpoint:         endpoint,
		ClientId:         clientId,
		ClientSecret:     clientSecret,
		JwtSecret:        jwtSecret,
		OrganizationName: organizationName,
	}
}

func GetUsers() ([]*User, error) {
	url := fmt.Sprintf("%s/api/get-users?owner=%s&clientId=%s&clientSecret=%s", authConfig.Endpoint, authConfig.OrganizationName, authConfig.ClientId, authConfig.ClientSecret)
	bytes, err := getBytes(url)
	if err != nil {
		return nil, err
	}

	var users []*User
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(name string) (*User, error) {
	url := fmt.Sprintf("%s/api/get-user?id=%s/%s&clientId=%s&clientSecret=%s", authConfig.Endpoint, authConfig.OrganizationName, name, authConfig.ClientId, authConfig.ClientSecret)
	bytes, err := getBytes(url)
	if err != nil {
		return nil, err
	}

	var user *User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user User) (bool, error) {
	return modifyUser("update-user", user)
}

func AddUser(user User) (bool, error) {
	return modifyUser("add-user", user)
}

func DeleteUser(user User) (bool, error) {
	return modifyUser("delete-user", user)
}
