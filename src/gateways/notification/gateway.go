package notification

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/postech-soat2-grupo16/clientes-api/entities"
	"os"
)

type Gateway struct {
	notificationTopic string
	notification      *sns.SNS
}

func NewGateway(notificationClient *sns.SNS) *Gateway {
	if notificationClient == nil {
		return NewGatewayMock()
	}
	return &Gateway{
		notificationTopic: os.Getenv("NOTIFICATION_TOPIC"),
		notification:      notificationClient,
	}
}

func (g *Gateway) ClientSubscriber(cliente *entities.Cliente) error {

	topicArn := g.notificationTopic
	protocol := "email"
	endpoint := cliente.Email

	filterPolicy := map[string]interface{}{
		"target": []string{cliente.Email},
	}
	filterPolicyJSON, err := json.Marshal(filterPolicy)
	if err != nil {
		fmt.Println("Error marshaling filter policy:", err)
		return err
	}

	fmt.Printf("Criando params para o cliente %d\n", cliente.ID)
	params := &sns.SubscribeInput{
		TopicArn:   aws.String(topicArn),
		Protocol:   aws.String(protocol),
		Endpoint:   aws.String(endpoint),
		Attributes: map[string]*string{"FilterPolicy": aws.String(string(filterPolicyJSON))},
	}

	fmt.Printf("Enviando Subscription de pagamento para %d\n", cliente.ID)
	resp, err := g.notification.Subscribe(params)
	if err != nil {
		fmt.Println("Erro ao enviar mensagem para a fila:", err)
		return nil
	}

	fmt.Printf("Subscription de pagamento para %d enviada com sucesso: %v\n", cliente.ID, resp.String())

	return nil
}

func NewGatewayMock() *Gateway {
	return nil
}

type GatewayMock struct {
}

func (g GatewayMock) ClientSubscriber(cliente *entities.Cliente) error {
	return nil
}
