package scenario

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetAllScenariosFromDir(scenariosPath string) (allScenarios []Scenario) {
	cars := getCarIdList(scenariosPath)

	for _, car := range cars {
		allScenarios = append(allScenarios, getScenario(car, scenariosPath))
	}
	return allScenarios
}

func getCarIdList(pathToScenarioFolder string) []string {
	var cars []string
	cars, _ = getListsOfDirsAndFiles(pathToScenarioFolder)
	if len(cars) == 0 {
		panic(fmt.Sprintf("No subdirectories for cars found in %v", pathToScenarioFolder))
	}
	log.Printf("[INFO] Found cars: %v\n", strings.Join(cars, ", "))
	return cars
}

func getScenario(carId, scenarioPath string) Scenario {
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
