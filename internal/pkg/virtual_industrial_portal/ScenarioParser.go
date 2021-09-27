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

func GetScenario(topic, scenarioPath string, loop bool) *Scenario{
	dirs, files := getListsOfDirsAndFiles(scenarioPath + "/" + topic)
	var scenarioStruct ScenarioStruct
	if(len(dirs) != 0){
		panic(fmt.Sprintf("[%v] Scenario folder contains dir: %v\n", topic, dirs))
	}

	if(len(files) > 1){
		log.Printf("[WARNING] [%v] Multiple scenario files have been found, only first file will be run", topic)
	}

	for _, file := range files{
		filePath := scenarioPath + "/" + topic + "/" + file

		matched, err := filepath.Match("*.json", filepath.Base(filePath))
		if err != nil {
            panic(fmt.Sprintf("[%v] Failed to match filename %v", topic, err.Error()))
        } else if matched {	
			scenarioStruct = parseJson(filePath)
			break
        }else{
			log.Printf("[WARNING] [%v] %v is not json file, ignoring\n", topic, filePath)
		}
	}


	log.Printf("[INFO] [%v] Found scenario files %v, creating scenario %v %v\n",topic, files, scenarioStruct.Map, scenarioStruct.Missions);
	scenario := NewScenario(scenarioStruct, topic, loop)
	return scenario
}

func parseJson(path string)(scenarioStruct ScenarioStruct){
	
	file, _ := ioutil.ReadFile(path)
	err := json.Unmarshal([]byte(file), &scenarioStruct)

	if(err != nil){
		panic(fmt.Sprintf("Unable to parse json file: %v error: %v", path, err.Error()))
	}

	return scenarioStruct
}

func getListsOfDirsAndFiles(path string)(dirs, files []string){
	//var dirsList, files []string
    dirInfo, err := ioutil.ReadDir(path)
    if err != nil {
        panic(fmt.Sprintf("Unable to access %v", path))
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