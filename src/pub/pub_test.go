package main

import (
	"encoding/json"
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// Teste de conexão com o broker MQTT
func TestConnectToMQTTBroker(t *testing.T) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("test-client")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Errorf("Falha ao conectar ao broker MQTT: %v", token.Error())
	} else {
		t.Log("Conexão com o broker MQTT estabelecida com sucesso.")
	}
}

// Teste de validação dos dados gerados
func TestGenerateData(t *testing.T) {
	data := GenerateData()

	// Verifique se todos os campos esperados estão presentes nos dados gerados
	expectedFields := []string{"alert", "id", "temperatura", "timestamp", "tipo"}
	for _, field := range expectedFields {
		if _, ok := data[field]; !ok {
			t.Errorf("Campo esperado %s não encontrado nos dados gerados", field)
			return
		}
	}
	t.Log("Dados gerados validados com sucesso.")
}

// Teste de confirmação de recebimento das publicações
func TestPublishAndReceiveMessages(t *testing.T) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("test-client")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Falha ao conectar ao broker MQTT: %v", token.Error())
	}

	// Inicia um subscriber MQTT para receber as publicações
	received := make(chan bool)
	token := client.Subscribe("/sensors", 0, func(client MQTT.Client, msg MQTT.Message) {
		// Verifique se a mensagem recebida é válida
		var data map[string]interface{}
		if err := json.Unmarshal(msg.Payload(), &data); err != nil {
			t.Errorf("Erro ao decodificar a mensagem JSON: %v", err)
			return
		}

		// Verifique se todos os campos esperados estão presentes na mensagem recebida
		expectedFields := []string{"alert", "id", "temperatura", "timestamp", "tipo"}
		for _, field := range expectedFields {
			if _, ok := data[field]; !ok {
				t.Errorf("Campo esperado %s não encontrado na mensagem recebida", field)
				return
			}
		}

		// Marque como recebido
		received <- true
	})
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Falha ao se inscrever no tópico MQTT: %v", token.Error())
	}

	// Adicione um atraso de 1 segundo para garantir que o subscriber tenha tempo de se conectar e se inscrever
	time.Sleep(1 * time.Second)

	// Publica uma mensagem com dados gerados
	data := GenerateData()
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Erro ao converter dados para JSON: %v", err)
	}

	msg := string(jsonData)
	token = client.Publish("/sensors", 0, false, msg)
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Falha ao publicar mensagem MQTT: %v", token.Error())
	}

	// Aguarda a confirmação de recebimento
	select {
	case <-received:
		// Mensagem recebida com sucesso
		t.Log("Mensagem recebida com sucesso.")
	case <-time.After(10 * time.Second):
		// Timeout - nenhum sinal de recebimento
		t.Fatalf("Timeout: Nenhuma mensagem recebida após 10 segundos")
	}
}