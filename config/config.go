package config

type Server struct {
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// auto
	AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
}
