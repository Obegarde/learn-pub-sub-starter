package pubsub
import(
	amqp "github.com/rabbitmq/amqp091-go"
	"encoding/json"
	"fmt"
	"context"
)
func PublishJSON[T any](ch *amqp.Channel, exchange,key string, val T) error{
	marshalledVal,err := json.Marshal(val)
	if err != nil{
		fmt.Println("Failed to marshal json")
		return err
	}
	msgToPublish := amqp.Publishing{
			ContentType: "application/json",
			Body: marshalledVal,
	}
	err = ch.PublishWithContext(context.Background(),exchange,key,false,false,msgToPublish)	
	if err != nil{
		fmt.Println("failed to publish to channel")
		return err
	}
	return nil
}

