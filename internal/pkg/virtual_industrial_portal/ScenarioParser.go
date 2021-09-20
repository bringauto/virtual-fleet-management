package virtual_industrial_portal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func GetListOfTopics(pathToScenarioFolder string) []string{
	var topics []string
	companies, _ := getListsOfDirsAndFiles(pathToScenarioFolder)
	for _, company := range companies{
		places, _ := getListsOfDirsAndFiles(pathToScenarioFolder + "/" + company)
		for _, place := range places{
			cars, _ := getListsOfDirsAndFiles(pathToScenarioFolder + "/" + company + "/" + place)
			for _, car := range cars{
				topics = append(topics, company + "/" + place + "/" + car)
			}
		}
	}
	log.Printf("[INFO] Parsed topics: %v\n", topics)
    
	return topics
}

func GetScenario(topic, scenarioPath string) *Scenario{
	dirs, files := getListsOfDirsAndFiles(scenarioPath + "/" + topic)
	var scenarioStruct ScenarioStruct
	if(len(dirs) != 0){
		panic(fmt.Sprintf("Scenario folder for topic %v contains dir: %v\n", topic, dirs))
	}

	if(len(files) > 1){
		log.Printf("[WARNING] multiple scenario files have been found, only first file will be run")
	}

	for _, file := range files{
		filePath := scenarioPath + "/" + topic + "/" + file

		matched, err := filepath.Match("*.json", filepath.Base(filePath))
		if err != nil {
            panic("Failed to match filename " + err.Error())
        } else if matched {	
			scenarioStruct = parseJson(filePath)
			break
        }else{
			log.Printf("[WARNING] %v is not json file, ignoring\n", filePath)
		}
	}


	log.Printf("[INFO] Found scenario files %v for %v, creating scenario %v %v\n", files, topic, scenarioStruct.Map, scenarioStruct.Missions);
	scenario := NewScenario(scenarioStruct)
	return scenario
}

func parseJson(path string)(scenarioStruct ScenarioStruct){
	
	file, _ := ioutil.ReadFile(path)
	err := json.Unmarshal([]byte(file), &scenarioStruct)

	if(err != nil){
		panic("Unable to parse json file: " + path + " error: " + err.Error())
	}

	return scenarioStruct
}

func getListsOfDirsAndFiles(path string)(dirs, files []string){
	//var dirsList, files []string
    dirInfo, err := ioutil.ReadDir(path)
    if err != nil {
        panic("Unable to access " + path)
    }
    for _, dir := range dirInfo {
		if(dir.IsDir()){
			dirs = append(dirs, dir.Name())
		}else{
			files = append(files, dir.Name())
		}
    } 
	return dirs, files;
}