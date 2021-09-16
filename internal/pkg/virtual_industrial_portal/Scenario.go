package virtual_industrial_portal

import(
	"log"
)
//"encoding/json"

type Scenario struct {
	missions	[] Mission

}

type Mission struct {
	fullStopList	[]string

}

func NewScenario(){
	log.Println("[INFO] creating new scenario")
}

func NewMission(){
	log.Println("[INFO] creating new mission")
}