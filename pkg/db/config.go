package db

type Mysql struct {
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	Protocol        string `yaml:"protocol"`
	Address         string `yaml:"address"`
	DBName          string `yaml:"dbName"`
	Driver          string `yaml:"driver"`
	ConnMaxLifeTime int    `yaml:"connMaxLifeTime"`
	MaxIdleConnNum  int    `yaml:"maxIdleConnNum"`
	MaxOpenConnNum  int    `yaml:"maxOpenConnNum"`
}

type Redis struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Password     string `yaml:"password"`
	DBIndex      int    `yaml:"db"`
	MinIdleConns int    `yaml:"minIdleConns"` //最小空闲连接数
	IdleTimeout  int    `yaml:"idleTimeout"`  //多久没有使用回收
	MaxConnAge   int    `yaml:"maxConnAge"`   //连接生命周期
	PoolSize     int    `yaml:"poolSize"`
}
