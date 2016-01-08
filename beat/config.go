package factbeat

type FactConfig struct {
	Period *int64
}

type ConfigSettings struct {
	Input FactConfig
}
