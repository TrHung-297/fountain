package env

type LogConfig struct {
	Path          string `json:"path,omitempty"`
	Prefix        string `json:"prefix,omitempty"`
	LogFileLevel  int    `json:"log_file_level,omitempty"`
	LogPrintLevel int    `json:"log_print_level,omitempty"`
	Description   string `json:"description,omitempty"`
}

var (
	LogConfigData *LogConfig
	LogPath       string
	LogPrefix     string
	LogFileLevel  int
	LogPrintLevel int
)
