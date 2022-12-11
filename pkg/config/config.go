package config

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

type ResolverConfig struct {
	PathToEmbeddings string
	PathToNames      string
}

func NewResolverConfig(embPath, namesPath string) *ResolverConfig {
	return &ResolverConfig{
		PathToEmbeddings: embPath,
		PathToNames:      namesPath,
	}
}
