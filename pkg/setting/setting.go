package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string
	PrefixUrl       string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	QrCodeSavePath string
	ExportSavePath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func Setup() {
	CfgPath := "conf/app.ini"
	Cfg, err := ini.Load(CfgPath)
	if err != nil {
		log.Fatalf("Fail to parse '%s' err:%v", CfgPath, err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err:%v", err)
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err:%v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err:%v", err)
	}
}
