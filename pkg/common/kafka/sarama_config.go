/**
 * Created by liemlhd on 30/Jan/2022
 */

package kafka

import (
	"time"

	"github.com/Shopify/sarama"
)

var DefaultSaramaConfig = defaultConfig()

func defaultConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.ClientID = "mengoaingu-golang"
	config.Version = sarama.V1_0_0_0
	// Consumer
	{
		// Always listen from first message
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
		config.Consumer.Return.Errors = true
	}
	// Sync producer
	{
		config.Producer.Return.Successes = true
		config.Producer.Retry.Max = 10
		config.Producer.Retry.Backoff = 2 * time.Second
	}
	return config
}
