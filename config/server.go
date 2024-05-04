package config

import "github.com/spf13/viper"

type server struct {
	port string
}

func (s *server) load() {
	viper.SetEnvPrefix("SERVER")
	s.port = viper.GetString("PORT")
}
