package virtual_industrial_portal

import (
	"fmt"
	"io/ioutil"
	"log"
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
	log.Printf("[INFO] found topics: %v\n", topics)
    
	return topics
}

func GetScenario(topic, scenarioPath string) *Scenario{
	dirs, files := getListsOfDirsAndFiles(scenarioPath + "/" + topic)

	if(len(dirs) != 0){
		panic(fmt.Sprintf("Scenario folder for topic %v contains dir: %v\n", topic, dirs))
	}

	//todo for each

	log.Printf("[INFO] Found scenario files %v for %v\n", files, topic);
	scenario := NewScenario()
	return scenario
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