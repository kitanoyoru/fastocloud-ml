package config

import "github.com/spf13/viper"

const (
	configDir = "./config"
)

func ReadConfigFile() (*FastocloudConfig, *FaceRecognitionResolverConfig) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(configDir)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("err") // TODO: Handle error
		}
	}

	fc := NewFastocloudConfig(
		viper.GetString("fastocloud.endpoint"),
		viper.GetString("fastocloud.login"),
		viper.GetString("fastocloud.password"),
	)

	rc := NewFaceRecognitionResolverConfig(
		viper.GetString("face-recognition.embeddingsPath"),
		viper.GetString("face-recognition.namesPath"),
	)

	return fc, rc
}
