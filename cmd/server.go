/*
Copyright © 2023 NAME HERE trongdt.working@gmail.com

*/
package cmd

import (
	"backend/pkg"
	"backend/pkg/bindings"
	"backend/pkg/common"
	"backend/pkg/common/kafka"
	"backend/pkg/config"
	"backend/pkg/firebase"
	"backend/pkg/middleware"
	gen "backend/proto/gen"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	wkafka "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var (
	cfgFile string
	cfgDir  string
)

// serverCmd represents the server2 command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			fx.Provide(
				func() context.Context {
					return context.Background()
				},
			),
			fx.Invoke(initConfig),
			fx.Provide(middleware.NewAuthClient),
			fx.Provide(firebase.SetupFirebase),
			fx.Provide(firebase.SetupJWT),
			fx.Provide(pkg.New),
			fx.Provide(config.NewConfig),
			bindings.ServicesOptions,
			fx.Provide(http.NewServeMux),
			fx.Invoke(
				startServer,
				// invokeMessaging,
			),
		)
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		workdir, err := os.Getwd()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".server" (without extension).
		viper.AddConfigPath(workdir)
		viper.SetConfigType("env")
		viper.SetConfigName("config.env")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func startServer(ctx context.Context, cfg *config.Config, app *pkg.App) {
	go startServerHTTP(ctx, cfg, app)
	go startServerGRPC(cfg, app)
}

func startServerGRPC(cfg *config.Config, handler *pkg.App) {
	address := fmt.Sprintf("%s:%d", cfg.HOST, cfg.GRPC_PORT)
	network := "tcp"
	l, err := net.Listen(network, address)
	if err != nil {
		log.Error("failed to listen to address ", err, " network ", network, " address ", address)
		// cancel()
	}

	log.Info("🌏 start server...", " address ", address)

	defer func() {
		if err1 := l.Close(); err != nil {
			log.Error("failed to close", err1, "network", network, "address", address)
		}
	}()

	err = handler.GrpcServer.Serve(l)
	if err != nil {
		log.Error("failed start gRPC server", err, "network", network, "address", address)
	}
}

func startServerHTTP(ctx context.Context, cfg *config.Config, handler *pkg.App) {
	gen.RegisterProfileServiceHandlerFromEndpoint(ctx, handler.Mux, "localhost:9090", []grpc.DialOption{grpc.WithInsecure()})
	gen.RegisterQuizzesServiceHandlerFromEndpoint(ctx, handler.Mux, "localhost:9090", []grpc.DialOption{grpc.WithInsecure()})
	gen.RegisterTaskServiceHandlerFromEndpoint(ctx, handler.Mux, "localhost:9090", []grpc.DialOption{grpc.WithInsecure()})
	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HOST, cfg.EXTERNAL_HTTP_PORT),
		Handler: cors(handler.Mux),
	}

	log.Info("start listening...", " address ", fmt.Sprintf("%s:%d", cfg.HOST, cfg.EXTERNAL_HTTP_PORT))

	if err := s.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		log.Error("failed to listen and serve", err)
	}
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func invokeMessaging() error {
	subscriber, err := wkafka.NewSubscriber(
		kafka.NewKafkaWatermillSubscriberConfig("trong_test_kafka"),
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return err
	}
	go consume(context.Background(), subscriber)

	publisher, err := wkafka.NewPublisher(
		kafka.NewKafkaWatermillPublisherConfig(),
		watermill.NewStdLogger(false, false),
	)
	go produce(publisher)
	if err != nil {
		return err
	}

	return err
}

func produce(pub *wkafka.Publisher) {
	// initialize a counter
	i := 0

	for {
		idgen := common.NewLocalSonyFlakeIdGenerator()
		id, err := idgen.NextID(context.Background())
		if err != nil {
			panic("could not generate id " + err.Error())
		}
		msg := message.NewMessage(id, []byte(fmt.Sprintf("Hello, world! %d", i)))
		err = pub.Publish("trong_test_kafka", msg)
		if err != nil {
			panic("could not send message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Millisecond * 5)
	}
}
func consume(ctx context.Context, sub *wkafka.Subscriber) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages

	// the `ReadMessage` method blocks until we receive the next event
	msg, err := sub.Subscribe(ctx, "trong_test_kafka")
	if err != nil {
		panic("could not read message " + err.Error())
	}
	// after receiving the message, log its value
	go process(msg)

}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		fmt.Printf("received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
