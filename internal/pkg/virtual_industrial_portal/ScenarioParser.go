package virtual_industrial_portal

import (
	"io/ioutil"
	"log"
)

func GetListOfTopics(pathToScenarioFolder string) []string{
	var topics []string
	companies := getListOfDirs(pathToScenarioFolder)
	for _, company := range companies{
		places := getListOfDirs(pathToScenarioFolder + "/" + company)
		for _, place := range places{
			cars := getListOfDirs(pathToScenarioFolder + "/" + company + "/" + place)
			for _, car := range cars{
				topics = append(topics, company + "/" + place + "/" + car)
			}
		}
	}
	log.Printf("[INFO] found topics: %v\n", topics)
    
	return topics
}

func GetScenario(topic, scenarioPath string) *Scenario{
	//todo implement
	scenario := NewScenario()
	return scenario
}

func getListOfDirs(path string)[]string{
	var dirs []string
    dirInfo, err := ioutil.ReadDir(path)
    if err != nil {
        panic("Unable to access " + path)
    }
    for _, dir := range dirInfo {
        dirs = append(dirs, dir.Name())
    } 
	return dirs
}