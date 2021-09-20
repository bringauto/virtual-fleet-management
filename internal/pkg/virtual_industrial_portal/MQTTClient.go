package virtual_industrial_portal

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/eclipse/paho.golang/paho"
)

type MQTTClient struct {
	client                     *paho.Client
	msgChan                    chan *paho.Publish
	server, username, password string
	tcpCon                     net.Conn
	vehicles				   []*Vehicle
}

var Client = MQTTClient{}

func (mqttClient *MQTTClient) Start(server, username, password string, scenariosPath string, loop bool) {
	log.Printf("[INFO] Connecting to broker at %v\n", server)

    defer func() {
        if r := recover(); r != nil {
            log.Printf("[ERROR] Error occured in mqqtClient start: %v\n", r)
        }
    }()

	mqttClient.server = server
	mqttClient.username = username
	mqttClient.password = password

	mqttClient.msgChan = make(chan *paho.Publish)


	topics := GetListOfTopics(scenariosPath)

	//todo parse json
	for _, topic := range topics{
		scenario := GetScenario(topic, scenariosPath, loop)
		mqttClient.vehicles = append(mqttClient.vehicles, NewVehicle(topic, scenario))
	}
	
	mqttClient.tcpConnect()
	mqttClient.mqttConnect()
	mqttClient.subscribe()
	go mqttClient.reconnectHandler()
	mqttClient.listen()
}

func (mqttClient *MQTTClient) tcpConnect() {
	var err error
	retry := time.NewTicker(1 * time.Second)
RetryLoop:
	for {
		select {
		case <-retry.C:
			mqttClient.tcpCon, err = net.Dial("tcp", mqttClient.server)

			if err != nil {
				log.Printf("[ERROR] Failed to connect to %s: %s\n", mqttClient.server, err)
			} else {
				retry.Stop()
				break RetryLoop
			}
		}
	}
}

func (mqttClient *MQTTClient) mqttConnect() {
	config := paho.ClientConfig{
		Router: paho.NewSingleHandlerRouter(func(m *paho.Publish) {
			mqttClient.msgChan <- m
		})}

	mqttClient.client = paho.NewClient(config)

	var expiry = uint32(20)

	var connectProerties = paho.ConnectProperties{
		SessionExpiryInterval: &expiry,
	}

	cp := &paho.Connect{
		KeepAlive:  30,
		CleanStart: true,
		Username:   mqttClient.username,
		Password:   []byte(mqttClient.password),
		Properties: &connectProerties,
	}

	if mqttClient.username != "" {
		cp.UsernameFlag = true
	}
	if mqttClient.password != "" {
		cp.PasswordFlag = true
	}
	mqttClient.client.Conn = mqttClient.tcpCon

	retry := time.NewTicker(1 * time.Second)
RetryLoop:
	for {
		select {
		case <-retry.C:
			ca, err := mqttClient.client.Connect(context.Background(), cp)

			if err != nil || ca.ReasonCode != 0 {
				log.Printf("[ERROR] Failed to connect to %s: %s\n", mqttClient.server, err)
			} else {
				retry.Stop()
				break RetryLoop
			}
		}
	}

	log.Println("[INFO] Connected to mqtt")
}

func (mqttClient *MQTTClient) reconnectHandler() {
	retry := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-retry.C:
			var message []byte
			success := mqttClient.publish("conn/test", message)
			if !success {
				for _, vehicle := range mqttClient.vehicles {
					vehicle.resetVehicle()
				}
				mqttClient.tcpConnect()
				mqttClient.mqttConnect()
				mqttClient.subscribe()
			}
		}
	}
}

func (mqttClient *MQTTClient) listen() {
	for m := range mqttClient.msgChan {
		for _, vehicle := range mqttClient.vehicles {
			if m.Topic == vehicle.daemonTopic {
				vehicle.parseMessage(m.Payload)
				continue
			}
		}
	}
	log.Println("[INFO] left listen")
}

func (mqttClient *MQTTClient) Disconnect() {
	if mqttClient.client != nil {
		d := &paho.Disconnect{ReasonCode: 0}
		mqttClient.client.Disconnect(d)
	}
	log.Println("[INFO] disconnecting from broker")
}

func (mqttClient *MQTTClient) subscribe() {
	var qos = 2
	for _, vehicle := range mqttClient.vehicles {
		daemonTopic := vehicle.daemonTopic
		sa, err := mqttClient.client.Subscribe(context.Background(), &paho.Subscribe{
			Subscriptions: map[string]paho.SubscribeOptions{
				daemonTopic: {QoS: byte(qos)},
			},
		})
		if err != nil {
			log.Printf("Failed to subscribe to %s : %d\n", daemonTopic, err)
		}
		if sa.Reasons[0] != byte(qos) {
			log.Printf("Failed to subscribe to %s : %d\n", daemonTopic, sa.Reasons[0])
		}
		log.Printf("[INFO] Subscribed to topic %s\n", daemonTopic)
	}
}

func (mqttClient *MQTTClient) publish(topic string, binaryMessage []byte) bool {
	_, err := mqttClient.client.Publish(context.Background(), &paho.Publish{
		Topic:   topic,
		QoS:     byte(2),
		Retain:  false,
		Payload: []byte(binaryMessage),
	})

	if err != nil {
		log.Printf("[ERROR] error sending message:%v\n", err)
		return false
	}
	return true
}
