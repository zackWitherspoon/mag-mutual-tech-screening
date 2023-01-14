package main

import (
	"encoding/csv"
	"log"
	"mag_matual_tech_project/constants"
	"mag_matual_tech_project/user"
	"mag_matual_tech_project/user/model"
	"os"
	"strconv"
	"time"
)

const (
	userCSV = "./UserInformation.csv"
)

type Configuration struct {
	Service user.Service
}

func (csvConfig Configuration) NewConfiguration() *Configuration {
	users := csvConfig.loadUsers(userCSV)
	return &Configuration{
		Service: user.NewService(users),
	}
}

func (csvConfig Configuration) loadUsers(filePath string) *[]model.User {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Cannot open CSV file:", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}

	var users []model.User
	for count, row := range rows {
		if count != 0 {
			id, _ := strconv.ParseInt(row[0], 0, 0)
			date, _ := time.Parse(constants.DateFormat, row[5])
			newUser := model.User{
				Id:          int(id),
				FirstName:   row[1],
				LastName:    row[2],
				Email:       row[3],
				Profession:  row[4],
				DateCreated: date,
				Country:     row[6],
				City:        row[7],
			}
			users = append(users, newUser)
		}
	}

	return &users
}
