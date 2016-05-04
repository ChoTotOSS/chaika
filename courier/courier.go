package courier

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
)

type Courier interface {
	Send(serviceName string, catalog string, level string, message string)
}

type LogInfo struct {
	Host string
	Port int64
	Type string
}

var couriers map[string]Courier
var expired map[string]int64
var kv *api.KV

func Setup(consulHost string, consulPort int64) {

	expired = make(map[string]int64)
	couriers = make(map[string]Courier)
	fmt.Println("Initilize couriers")

	client, err := api.NewClient(&api.Config{
		Address: consulHost + ":" + strconv.FormatInt(consulPort, 10),
	})
	//api.DefaultConfig())
	CheckError(err)
	// Get a handle to the KV API
	kv = client.KV()
}

func Get(serviceName string) Courier {

	now := time.Now().Unix()

	if expiredTime, ok := expired[serviceName]; ok && expiredTime > now {
		return couriers[serviceName]
	}

	config := GetLogOutput(serviceName)

	couriers[serviceName] = CreateGelf(serviceName, config.Host, config.Port)
	expired[serviceName] = now + 5
	return couriers[serviceName]
}

func GetLogOutput(serviceName string) LogInfo {

	config := LogInfo{
		Host: "10.50.10.3",
		Port: 12201,
		Type: "gelf",
	}

	hostPair, _, err := kv.Get(serviceName+"/log/host", nil)

	CheckError(err)

	if hostPair != nil {
		config.Host = string(hostPair.Value)
	}

	portPair, _, err := kv.Get(serviceName+"/log/port", nil)
	CheckError(err)

	if hostPair != nil {
		config.Port, _ = strconv.ParseInt(string(portPair.Value), 10, 64)
	}

	typePair, _, err := kv.Get(serviceName+"/log/type", nil)
	CheckError(err)

	if typePair != nil {
		config.Type = string(typePair.Value)
	}

	return config
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(3)
	}
}
