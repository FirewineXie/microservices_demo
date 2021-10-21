package th_nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var (
	clientConfig  constant.ClientConfig
	serverConfigs []constant.ServerConfig
	err           error
	iClient       naming_client.INamingClient
)

func CreateClientConfig() {
	clientConfig = constant.ClientConfig{
		NamespaceId:         "1cffcc09-be70-4d6e-922b-3598de9b0df8",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}
}

func CreateMultipleServerConfig() {
	serverConfigs = []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

}

// CreateConfigClient 创建服务发现客户端
func CreateConfigClient() (naming_client.INamingClient, error) {
	CreateClientConfig()
	CreateMultipleServerConfig()
	iClient, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	return iClient, err
}

// RegisterInstance 服务注册
func RegisterInstance() (bool, error) {
	return iClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "10.0.0.11",
		Port:        8848,
		ServiceName: "demo.go",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "cluster-a", // 默认值DEFAULT
		GroupName:   "group-a",   // 默认值DEFAULT_GROUP
	})
}

// DeregisterInstance 注销实例
func DeregisterInstance() (bool, error) {
	return iClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          "10.0.0.11",
		Port:        8848,
		ServiceName: "demo.go",
		Ephemeral:   true,
		Cluster:     "cluster-a", // 默认值DEFAULT
		GroupName:   "group-a",   // 默认值DEFAULT_GROUP
	})
}

// SelectOneHealthyInstance 加权随机轮询获取一个健康的实例
func SelectOneHealthyInstance() (*model.Instance, error) {
	return iClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",             // 默认值DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
	})

}
