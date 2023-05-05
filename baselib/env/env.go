package env

type EnvConfig struct {
	Environment    string `json:"environment,omitempty"`
	Addr           string `json:"addr,omitempty"`
	DCName         string `json:"dc_name,omitempty"`
	HostName       string `json:"host_name,omitempty"`
	PodName        string `json:"pod_name,omitempty"`
	PodID          string `json:"pod_id,omitempty"`
	ServiceName    string `json:"service_name,omitempty"`
	EndpointPrefix string `json:"endpoint_prefix,omitempty"`
}

var (
	EnvConfigData  *EnvConfig
	Environment    string
	Addr           string
	DCName         string
	HostName       string
	PodName        string
	PodID          string
	ServiceName    string
	EndpointPrefix string
)
