package config

import (
	"log"
	"os"

	"github.com/caarlos0/env"
)

const KAFKA_BROKE_EVENT = "192.168.2.19:9092"
const KAFKA_BROKE_CAPTURE = "192.168.2.19:9092"

const KAFKA_PARTITION_EVENT = 0
const KAFKA_PARTITION_CAPTURE = 0

// OffsetNewest int64 = -1
// OffsetOldest int64 = -2
const KAFKA_OFFSET_EVENT = -1
const KAFKA_OFFSET_CAPTURE = -1

const KAFKA_TOPIC_EVENT = "face-event-topic"
const KAFKA_TOPIC_CAPTURE = "face-captured-topic"

type ArceeConfig struct {
	Host          string
	TimeoutSecond int
	ConnectNumber int
}

type ServiceConfig struct {
	IsOpen     bool
	ConfigType uint
	ConfigInfo interface{}
}

type FaceDeep struct {
	ImageSaveServiceConfig  *ServiceConfig
	BACKEND_API             string  `json:"thor" env:"BACKEND_API,required"`
	KAFKA_BROKE_EVENT       string  `json:"event_broke" env:"KAFKA_BROKE_EVENT,required"`
	KAFKA_BROKE_CAPTURE     string  `json:"capture_broke" env:"KAFKA_BROKE_CAPTURE,required"`
	KAFKA_PARTITION_EVENT   int     `json:"event_partition" env:"KAFKA_PARTITION_EVENT" envDefault:"0"`
	KAFKA_PARTITION_CAPTURE int     `json:"capture_partition" env:"KAFKA_PARTITION_CAPTURE" envDefault:"0"`
	KAFKA_OFFSET_EVENT      int     `json:"event_offset" env:"KAFKA_OFFSET_EVENT" envDefault:"-1"`
	KAFKA_OFFSET_CAPTURE    int     `json:"capture_offset" env:"KAFKA_OFFSET_CAPTURE" envDefault:"-1"`
	KAFKA_TOPIC_EVENT       string  `json:"face_event_topic" env:"KAFKA_TOPIC_EVENT" envDefault:"warned-push"`
	KAFKA_TOPIC_CAPTURE     string  `json:"face_captured_topic" env:"KAFKA_TOPIC_CAPTURE" envDefault:"captured-push"`
	KAFKA_TOPIC_WHITE       string  `json:"face_white_topic" env:"KAFKA_TOPIC_WHITE" envDefault:"white-push"`
	HTTP_PORT               string  `json:"http_port" env:"HTTP_PORT" envDefault:":8082"`
	KAFKA_TOPIC_SENDER      string  `json:"sender_topic" env:"KAFKA_TOPIC_SENDER" envDefault:"facedeplin"`
	KAFKA_BROKE_SENDER      string  `json:"sender_broke" env:"KAFKA_BROKE_SENDER,required"`
	THRESHOLD               float32 `json:"threshold" env:"THRESHOLD,required" envDefault:"0.85"`
	DB_ENGINE               string  `json:"db_engine" env:"DB_ENGINE,required" envDefault:"postgres"`
	DB_CONN                 string  `json:"db_conn" env:"DB_CONN,required"`
	OLYMPUS_ADDR            string  `json:"olympus_addr" env:"OLYMPUS_ADDR" envDefault:" "`
	LOKI_API                string  `json:"loki" env:"LOKI_API" envDefault:" "`
}

func GetConfig() FaceDeep {
	f := FaceDeep{}
	err := env.Parse(&f)
	if err != nil {
		log.Fatalf("[CONFIG] parse env config error, %s", err.Error())
		os.Exit(-1)
	}
	return f
}
