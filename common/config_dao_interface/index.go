package config_dao_interface

type ConfigDao interface {
	Get(key string) string
	Set(key, value string) error
	Del(key string) error
}
