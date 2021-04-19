package config

type Server struct {
	Mysql     Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Sqlite    Sqlite    `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Casbin    Casbin    `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Captcha   Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Log       Log       `mapstructure:"log" json:"log" yaml:"log"`
}

type System struct {
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" yaml:"model-path"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}


type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`
}

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile" yaml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	File    string `mapstructure:"file" json:"file" yaml:"file"`
}

type Sqlite struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
	LogMode  bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Wechat struct {
	Appid string `mapstructure:"appid" json:"appid" yaml:"appid"`
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
}

type Ipfs struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
}

type Fabsdk struct {
	ConfigPath string `mapstructure:"config-path" yaml:"config-path" json:"configPath"`
	UserName string `mapstructure:"user-name" yaml:"user-name" json:"userName"`
	ChannelName string `mapstructure:"channel-name" yaml:"channel-name" json:"channelName"`
	ChaincodeName string `mapstructure:"chaincode-name" yaml:"chaincode-name" json:"chaincodeName"`
	AdminPriHex string `mapstructure:"admin-priHex" yaml:"admin-priHex" json:"adminPriHex"`
	AdminPukHex string `mapstructure:"admin-pukHex" yaml:"admin-pukHex" json:"adminPukHex"`
	StartNum int `mapstructure:"start-num" yaml:"start-num" json:"startNum"`
	HandlerName string `mapstructure:"handler-name" yaml:"handler-name" json:"handlerName"`
}
