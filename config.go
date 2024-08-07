package database

import (
    "fmt"
    "net/url"
)

// Config 适用于单个连接的配置
type Config struct {
    Type        string `toml:"type" json:"type"`                   //数据库类型
    Host        string `toml:"host" json:"host"`                   //连接地址
    Port        int    `toml:"port" json:"port"`                   //端口
    User        string `toml:"user" json:"user"`                   //用户
    Pass        string `toml:"pass" json:"pass"`                   //密码
    Database    string `toml:"database" json:"database"`           //数据库名
    Param       string `toml:"param" json:"param"`                 //连接参数
    SlowLogTime int    `toml:"slow_log_time" json:"slow_log_time"` //慢日志阈值（毫秒）0为不开启
}

// ConfigMap 适用于多个连接的配置
type ConfigMap map[string]Config

// GenerateDsn 根据配置生成sdn连接信息
func (conf Config) GenerateDsn() string {
    dsn := ""
    dsnParam := ""

    if conf.Type == "mysql" {
        if conf.Param != "" {
            dsnParam = "?" + conf.Param
        }
        dsnF := "%s:%s@(%s:%d)/%s%s"
        dsn = fmt.Sprintf(dsnF, conf.User, url.QueryEscape(conf.Pass), conf.Host, conf.Port, conf.Database, dsnParam)
    } else if conf.Type == "sqlserver" {
        if conf.Param != "" {
            dsnParam = "&" + conf.Param
        }
        dsnF := "sqlserver://%s:%s@%s:%d?database=%s%s"
        dsn = fmt.Sprintf(dsnF, conf.User, url.QueryEscape(conf.Pass), conf.Host, conf.Port, conf.Database, dsnParam)
    } else if conf.Type == "postgres" {
        if conf.Param != "" {
            dsnParam = "?" + conf.Param
        }
        dsnF := "host=%s user=%s password=%s dbname=%s port=%d %s"
        dsn = fmt.Sprintf(dsnF, conf.Host, conf.User, url.QueryEscape(conf.Pass), conf.Database, conf.Port, dsnParam)
    } else if conf.Type == "oracle" {
        dsnF := "%s/%s@%s:%d/%s"
        dsn = fmt.Sprintf(dsnF, conf.User, url.QueryEscape(conf.Pass), conf.Host, conf.Port, conf.Database)
    }

    return dsn
}
