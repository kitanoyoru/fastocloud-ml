package config

type CliConfig struct {
	Mode int
}

func NewCliConfig(mode int) *CliConfig {
	return &CliConfig{
		Mode: mode,
	}
}

type FastocloudConfig struct {
	Endpoint string
	Login    string
	Password string
}

func NewFastocloudConfig(endpoint, login, password string) *FastocloudConfig {
	return &FastocloudConfig{
		Endpoint: endpoint,
		Login:    login,
		Password: password,
	}
}

type FaceRecognitionResolverConfig struct {
	PathToEmbeddings string
	PathToNames      string
}

func NewFaceRecognitionResolverConfig(embPath, namesPath string) *FaceRecognitionResolverConfig {
	return &FaceRecognitionResolverConfig{
		PathToEmbeddings: embPath,
		PathToNames:      namesPath,
	}
}
