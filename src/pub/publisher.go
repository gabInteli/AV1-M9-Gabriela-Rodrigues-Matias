package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	time "time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)


// Função para gerar dados aleatórios de sensores
func GenerateData() map[string]interface{} {
	items := []string{
		"Geladeira",
		"Freezer",
	}
	min := -50
	max := 50

	data := map[string]interface{}{
		"id":             rand.Intn(1000),
		"tipo":           items[rand.Intn(len(items))],
		"temperatura":   rand.Intn(max-min+1) + min,
		"timestamp":     time.Now().Format(time.RFC3339),
	}
	alert, alertMessage := isAlert(data["tipo"].(string), data["temperatura"].(int))
	data["alert"] = alertMessage
	if alertMessage != "" {
		fmt.Println(alert)
	}
	return data
}

func isAlert(tipo string, temperature int) (bool, string) {
	alertMessage := ""
	if tipo == "Freezer" {
		if temperature > -15 {
			alertMessage = "ALERTA: Temperatura ALTA"
			return true, alertMessage
		} else if temperature < -25 {
			alertMessage = "ALERTA: Temperatura BAIXA"
			return true, alertMessage
		}
	} else if tipo == "Geladeira" {
		if temperature > 10 {
			alertMessage = "ALERTA: Temperatura ALTA"
			return true, alertMessage
		} else if temperature < 2 {
			alertMessage = "ALERTA: Temperatura BAIXA"
			return true, alertMessage
		}
	}
	return false, alertMessage
}


func main() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		data := GenerateData()

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting data to JSON", err)
			return
		}

		msg := string(jsonData)

		token := client.Publish("/sensors", 1, false, msg) // QoS 1
		token.Wait()

		fmt.Println("Published:", msg)
		time.Sleep(2 * time.Second)
	}
}
