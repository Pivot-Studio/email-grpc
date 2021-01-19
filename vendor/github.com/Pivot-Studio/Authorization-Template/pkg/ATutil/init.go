package ATutil

var hostname string

func init() {
	conf := ReadSettingsFromFile("Config.json")
	aes_key = conf.Key
	hostname = conf.HostName
}
