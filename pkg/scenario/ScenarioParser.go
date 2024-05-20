package scenario

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// TODO func find scenarios in directory?

func GetCarIdList(pathToScenarioFolder string) []string {
	var cars []string
	cars, _ = getListsOfDirsAndFiles(pathToScenarioFolder)
	log.Printf("[INFO] Found cars: %v\n", cars)

	return cars
}

func GetScenario(carId, scenarioPath string) Scenario {
	dirs, files := getListsOfDirsAndFiles(scenarioPath + "/" + carId)
	var scenarioStruct ScenarioStruct
	if len(dirs) != 0 {
		panic(fmt.Sprintf("[%v] Scenario folder contains dir: %v\n", carId, dirs))
	}

	if len(files) > 1 {
		log.Printf("[WARNING] [%v] Multiple scenario files have been found, only first file will be run", carId)
	}

	for _, file := range files {
		filePath := scenarioPath + "/" + carId + "/" + file

		matched, err := filepath.Match("*.json", filepath.Base(filePath))
		if err != nil {
			panic(fmt.Sprintf("[%v] Failed to match filename %v", carId, err.Error()))
		} else if matched {
			scenarioStruct = parseJson(filePath)
			break
		} else {
			log.Printf("[WARNING] [%v] %v is not json file, ignoring\n", carId, filePath)
		}
	}

	//if scenarioStruct.Missions == nil {
	//	log.Printf("[WARNING] [%v] Found scenario files %v, don't contain missions\n", CarId, files)
	//	return nil
	//} else if scenarioStruct.Routes == nil {
	//	log.Printf("[WARNING] [%v] Found scenario files %v, don't contain routes\n", CarId, files)
	//	return nil
	//}
	scenario := NewScenario(scenarioStruct, carId)
	log.Printf("[INFO] [%v] Found scenario files %v, creating scenario %v Missions: %v Routes: %v\n", carId, files, scenarioStruct.Map, scenario.Missions, scenario.Routes)
	return scenario

}

func parseJson(path string) (scenarioStruct ScenarioStruct) {

	file, _ := os.ReadFile(path)
	err := json.Unmarshal([]byte(file), &scenarioStruct)

	if err != nil {
		panic(fmt.Sprintf("Unable to parse json file: %v error: %v", path, err.Error()))
	}

	return scenarioStruct
}

func getListsOfDirsAndFiles(path string) (dirs, files []string) {
	dirInfo, err := os.ReadDir(path)
	if err != nil {
		panic(fmt.Sprintf("Unable to access %v", path))
	}
	for _, dir := range dirInfo {
		if dir.IsDir() {
			dirs = append(dirs, dir.Name())
		} else {
			files = append(files, dir.Name())
		}
	}
	return dirs, files
}
