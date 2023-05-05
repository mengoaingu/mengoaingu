package kafka

import (
	wkafka "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/spf13/viper"
)

func getKafkaBrokersCfg() []string {
	brokers := []string{"localhost:9092"}
	v := viper.Get("services.kafka.brokers")
	if viper.IsSet("services.kafka.brokers") {
		brokers = v.([]string)
	}
	return brokers
}

func NewKafkaWatermillSubscriberConfig(groupName string) wkafka.SubscriberConfig {
	return wkafka.SubscriberConfig{
		Brokers:               getKafkaBrokersCfg(),
		Unmarshaler:           wkafka.DefaultMarshaler{},
		OverwriteSaramaConfig: DefaultSaramaConfig,
		ConsumerGroup:         groupName,
	}
}

func NewKafkaWatermillPublisherConfig() wkafka.PublisherConfig {
	return wkafka.PublisherConfig{
		Brokers:               getKafkaBrokersCfg(),
		Marshaler:             wkafka.DefaultMarshaler{},
		OverwriteSaramaConfig: DefaultSaramaConfig,
	}
}
