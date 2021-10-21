package pkg

import (
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"microservices_demo/third_party/th_nacos"
)

var (
	err     error
	iClient naming_client.INamingClient
)

const (
	clusters              = "microservice"
	serviceAd             = "service_ad"
	serviceCart           = "service_cart"
	serviceCheckout       = "service_checkout"
	serviceCurrency       = "service_currency"
	serviceEmail          = "service_email"
	servicePayment        = "service_payment"
	serviceProductCatalog = "service_product_catalog"
	serviceRecommendation = "service_recommendation"
	serviceShipping       = "service_shipping"
)

// ConnectNacos 获取nacos 的客户端
func ConnectNacos() {
	iClient, err = th_nacos.CreateConfigClient()
	if err != nil {
		panic(err)
	}
}

// RegisterInstance 服务注册
func RegisterInstance() (bool, error) {

	return iClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "0.0.0.0",
		Port:        9002,
		ServiceName: serviceCheckout,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: clusters,  // 默认值DEFAULT
		GroupName:   serviceCheckout, // 默认值DEFAULT_GROUP
	})
}

// DeregisterInstance 注销实例
func DeregisterInstance() (bool, error) {
	return iClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          "0.0.0.0",
		Port:        9002,
		ServiceName: serviceCheckout,
		Ephemeral:   true,
		Cluster:     clusters, // 默认值DEFAULT
		GroupName:   serviceCheckout,   // 默认值DEFAULT_GROUP
	})
}

// GetHealthServerAd 获取ad 服务健康实例
func GetHealthServerAd() (*model.Instance, error) {
	const name = "service_ad"
	const clusters = "microservice"
	return iClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
		GroupName:   name,               // 默认值DEFAULT_GROUP
		Clusters:    []string{clusters}, // 默认值DEFAULT
	})
}

func GetHealthServerCart() (*model.Instance, error) {
	const name = "service_cart"

	return iClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
		GroupName:   name,               // 默认值DEFAULT_GROUP
		Clusters:    []string{clusters}, // 默认值DEFAULT
	})
}
func GetHealthServerCheckout() (*model.Instance, error) {
	const name = "service_checkout"
	const clusters = "microservice"
	return iClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
		GroupName:   name,               // 默认值DEFAULT_GROUP
		Clusters:    []string{clusters}, // 默认值DEFAULT
	})
}
func GetHealthServerCurrency() (*model.Instance, error) {
	const name = "service_currency"
	const clusters = "microservice"
	return iClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
		GroupName:   name,               // 默认值DEFAULT_GROUP
		Clusters:    []string{clusters}, // 默认值DEFAULT
	})
}

func GetHealthServerEmail() (*model.Instance, error) {
	const name = "service_email"
	const clusters = "microservice"
	return iClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
		GroupName:   name,               // 默认值DEFAULT_GROUP
		Clusters:    []string{clusters}, // 默认值DEFAULT
	})
}

func GetHealthServerPayment() (*model.Instance, error) {
	const name = "service_payment"
	const clusters = "microservice"
	return iClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
		GroupName:   name,               // 默认值DEFAULT_GROUP
		Clusters:    []string{clusters}, // 默认值DEFAULT
	})
}

func GetHealthServerProductCatalog() (*model.Instance, error) {
	const name = "service_product_catalog"
	const clusters = "microservice"
	return iClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
		GroupName:   name,               // 默认值DEFAULT_GROUP
		Clusters:    []string{clusters}, // 默认值DEFAULT
	})
}

func GetHealthServerRecommendation() (*model.Instance, error) {
	const name = "service_recommendation"
	const clusters = "microservice"
	return iClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
		GroupName:   name,               // 默认值DEFAULT_GROUP
		Clusters:    []string{clusters}, // 默认值DEFAULT
	})
}

func GetHealthServerShipping() (*model.Instance, error) {
	const name = "service_shipping"
	const clusters = "microservice"
	return iClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
		GroupName:   name,               // 默认值DEFAULT_GROUP
		Clusters:    []string{clusters}, // 默认值DEFAULT
	})
}
