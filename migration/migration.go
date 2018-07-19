package migration

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/go-squads/genrevan-scheduler/utils"
)

var basepath = utils.GetRootFolderPath()

func RunMigration() error {
	err := model.SetupDatabase("testing")
	if err != nil {
		return err
	}

	queryString, err := GetStringFromFile(basepath + "/migration/schema.sql")

	if err != nil {
		return err
	}

	model.Db.Exec(*queryString)

	return nil
}

func RunSeeder() error {
	err := model.SetupDatabase("testing")
	if err != nil {
		return err
	}

	queryString, err := GetStringFromFile(basepath + "/migration/seeder.sql")

	if err != nil {
		return err
	}

	model.Db.Exec(*queryString)

	return nil
}

func GetStringFromFile(filename string) (*string, error) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, errors.New("File not found")
	}

	queryString := fmt.Sprintf("%s", content)

	return &queryString, nil
}
