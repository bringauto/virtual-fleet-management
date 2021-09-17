package virtual_industrial_portal

func GetListOfTopics(pathToScenarioFolder string) []string{
	//todo implement
	var topics []string
	topics = append(topics, "roboauto/kralovopolska/car1")
	topics = append(topics, "faulhorn/borsodchem/car1")
	topics = append(topics, "bringauto/default/car1")
	return topics
}

func GetScenario(topic, scenarioPath string) *Scenario{
	//todo implement
	scenario := NewScenario()
	return scenario
}