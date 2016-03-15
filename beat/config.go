package factbeat

type FactConfig struct {
	Period *int64
	Facter *string
}

type ConfigSettings struct {
	Input FactConfig
}
