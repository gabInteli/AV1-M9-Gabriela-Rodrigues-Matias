## Gabriela Rodrigues Matias - Avaliação 1
### Engenharia da Computação


# Simulador IoT - Medição de Temperatura: Mercado

## Simulação de Envio e Captação de Dados de Sensores
Criar um simulador de dispositivos IoT utilizando o protocolo MQTT através do uso da biblioteca Eclipse Paho. 
Este repositório inclui uma estrutura de publisher e subscriber desenvolvida com Go, que publica mensagens JSON em um tópico MQTT e um que recebe, utilizando o Mosquitto como broker. As mensagens contêm dados simulados de medições, respeitando especificações técnicas como faixa de medição e alcance espectral, alem disso para garantir uma variação de dados, os valores são aleatórios dentro de um intervalo estabelecido pelo código.

## Requisitos
- Mosquitto
- Paho MQTT
- Python
- Go

##  Modo de Execução 

### Configuração do Mosquitto
Você deve configurar o Mosquitto antes de iniciá-lo. Supondo que você tenha um arquivo de configuração chamado mosquitto.conf, você pode iniciar o Mosquitto com a seguinte linha de comando:

```
mosquitto -c mosquitto.conf
```

Esse comando inicia o broker Mosquitto com as configurações definidas em mosquitto.conf.

### Instalando Dependências - Go Mod
Acesse o diretorio que contem as dependências necessárias para cada função: 

Para o publisher:
```
/src/pub
```

Para o subscriber:
```
/src/sub
```

Acione as dependências para cada uma das pastas, com: 
```
go mod tidy
```

### Executando o Publisher

Para executar o script do publisher, navegue até o diretório onde o arquivo publisher.go: 
```
/src/pub
```

e execute o seguinte comando no terminal:
```
go run publisher.go
```

Isso iniciará o script que publicará mensagens JSON simuladas para o tópico MQTT no intervalo definido. 
O script imprimirá no terminal cada mensagem que for publicada. Como demonstrado na imagem abaixo:

![Publisher](./img/av1-pub.png)

### Executando o Subscriber

Para executar o script do subscriber, navegue até o diretório onde o arquivo subscriber.go: 
```
/src/sub
```

e execute o seguinte comando no terminal:
```
go run subscriber.go
```

Isso iniciará o script que receberá as mensagens JSON simuladas para o tópico MQTT no intervalo definido. O script imprimirá no terminal cada mensagem que for recebida. Como demonstrado na imagem abaixo:

![Subscriber](./img/p1-sub.png)

### Rodando os Testes

Acesse o diretorio: 

```
/src/pub
```

Rode o comando: 
```
go test -v
```

Resultado esperado: 
```
=== RUN   TestConnectToMQTTBroker
    pub_test.go:20: Conexão com o broker MQTT estabelecida com sucesso.
--- PASS: TestConnectToMQTTBroker (0.00s)
=== RUN   TestGenerateData
    pub_test.go:36: Dados gerados validados com sucesso.
--- PASS: TestGenerateData (0.00s)
=== RUN   TestPublishAndReceiveMessages
    pub_test.go:99: Mensagem recebida com sucesso.
--- PASS: TestPublishAndReceiveMessages (0.00s)
PASS
ok      paho-go 0.003s
```
______________________________________________________________________________________________

## Teste de Conexão com o Broker MQTT

### Propósito:
Este teste verifica se é possível estabelecer uma conexão com o broker MQTT especificado.

### Entrada:
- Nenhuma entrada explícita.

### Saída Esperada:
- Conexão bem-sucedida com o broker MQTT.

## Teste de Geração de Dados

### Propósito:
Este teste verifica se os dados gerados pela função `GenerateData` estão completos e contêm todos os campos esperados.

### Entrada:
- Nenhuma entrada explícita.

### Saída Esperada:
- Mapa contendo os seguintes campos:
 - "id": Identificação do Sensor
 - "tipo": Se é um sensor do Freezer ou da Geladeira
 - "temperatura": Medição de Temperatura do Sensor
  "timestamp": Data de registro

## Teste de Publicação e Recebimento de Mensagens MQTT

### Propósito:
Este teste simula o processo de publicação e recebimento de mensagens MQTT.

### Entrada:
- Mensagem JSON contendo dados de sensores gerados aleatoriamente.

### Saída Esperada:
- Mensagem JSON recebida com sucesso, contendo os mesmos campos e valores da mensagem publicada.

### Passos do Teste:
1. Publicar uma mensagem MQTT contendo dados de sensores gerados aleatoriamente.
2. Aguardar a recebimento da mensagem publicada no tópico MQTT especificado.
3. Verificar se a mensagem recebida contém todos os campos esperados e se os valores estão dentro dos limites esperados.
4. Reportar um erro se a mensagem não for recebida dentro do tempo limite especificado.

Para cada teste, são fornecidas explicações detalhadas sobre o que o teste faz e qual é o seu propósito.

### Demonstração: 

A demonstração pode ser verificada no vídeo abaixo:  
[Vídeo de Demonstração](https://www.youtube.com/embed/jS9fEmYpGwQ?si=EN-pg6KznCJKwJjy")

Obs.: Tive problemas com o teste final. Como apresentada a falha no vídeo. 