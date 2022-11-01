package config

type Config struct {
	Addr   string `json:"addr"`
	Port   uint16 `json:"port"`
	DbPath string `json:"db_path"`
}
