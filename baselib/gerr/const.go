package gerr

var (
	LogBackEndPrefix          string
	LogBackEndModulePrefix    string
	LogBackEndMainInfoPrefix  string
	LogBackEndMainErrorPrefix string
)

var (
	LogInfoPrefix  = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainInfoPrefix
	LogErrorPrefix = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainErrorPrefix
)

func SetLogBackEndPrefix(prefix string) {
	LogBackEndPrefix = prefix
	LogInfoPrefix = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainInfoPrefix
	LogErrorPrefix = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainErrorPrefix
}

func SetLogBackEndModulePrefix(prefix string) {
	LogBackEndModulePrefix = prefix
	LogInfoPrefix = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainInfoPrefix
	LogErrorPrefix = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainErrorPrefix
}

func SetLogBackEndMainInfoPrefix(prefix string) {
	LogBackEndMainInfoPrefix = prefix
	LogInfoPrefix = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainInfoPrefix
}

func SetLogBackEndMainErrorPrefix(prefix string) {
	LogBackEndMainErrorPrefix = prefix
	LogErrorPrefix = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainErrorPrefix
}
