package repository

import (
	"encoding/csv"
	"os"
	"strconv"
)

// Indexes for the city data list
const (
	nameIdx = iota
	regionIdx
	districtIdx
	populationIdx
	foundationIdx
)

// Error constants
const (
	errNotFoundId = "ERROR: CITY WITH THIS ID WAS NOT FOUND"
)

// DataBase contains a list of cities, and the last free identifier
type DataBase struct {
	records map[int][]string
	lastID  int
}

// NewDataBase creates a database.
// Opens a csv file and copies the data into the DataBase structure.
// If successful, it returns *DataBase
func NewDataBase(filePath string) (*DataBase, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	cities, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	cityList := new(DataBase)
	cityList.records = make(map[int][]string)

	if len(cities) == 0 {
		return cityList, nil
	}

	for _, city := range cities {
		id, _ := strconv.Atoi(city[0])
		if err != nil {
			return nil, err
		}
		cityList.records[id] = make([]string, 5)
		copy(cityList.records[id], city[1:])
	}

	cityList.lastID = 0
	for cityID := range cityList.records {
		if cityList.lastID < cityID {
			cityList.lastID = cityID
		}
	}
	return cityList, nil
}

// SaveCSV saves the database to a csv file.
// Creates a csv file and copies data from the DataBase structure into it.
// If successful, overwrites the csv file into the filePath file.
func (db *DataBase) SaveCSV(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var data [][]string
	for id, description := range db.records {
		var cityLine []string
		cityLine = append(cityLine, strconv.Itoa(id))
		cityLine = append(cityLine, description...)
		data = append(data, cityLine)
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll(data)
	if err != nil {
		return err
	}
	return nil
}
