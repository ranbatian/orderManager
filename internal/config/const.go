package config

type Config struct {
	Port     int      `yaml:"port"`
	Database Database `yaml:"database"`
	Log      Log      `yaml:"log"`
}

type Log struct {
	File string `yaml:"file"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Hostname string `yaml:"hostname"`
	DBName   string `yaml:"dbName"`
}
