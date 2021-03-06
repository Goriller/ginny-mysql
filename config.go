package mysql

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Config DB基础配置
type Config struct {
	WDB         Source   `json:"wdb" mapstructure:"wdb"`
	RDBs        []Source `json:"rdbs" mapstructure:"rdbs"`
	DBName      string   `json:"db_name" mapstructure:"db_name"`
	MaxOpenConn int      `json:"max_open_conn" mapstructure:"max_open_conn"`
	MaxIdleConn int      `json:"max_idle_conn" mapstructure:"max_idle_conn"`
	MaxLifetime int      `json:"max_lifetime" mapstructure:"max_lifetime"`
	Keepalive   int      `json:"keepalive" mapstructure:"keepalive"`
}

// Source DB部署实例数据源配置
type Source struct {
	Host string `json:"host" mapstructure:"host"`
	User string `json:"user" mapstructure:"user"`
	Pass string `json:"pass" mapstructure:"pass"`
}

// NewConfig
func NewConfig(v *viper.Viper) (*Config, error) {
	var err error
	o := new(Config)
	if err = v.UnmarshalKey("mysql", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}

	if o.RDBs == nil || len(o.RDBs) == 0 {
		o.RDBs = []Source{
			{
				Host: o.WDB.Host,
				User: o.WDB.User,
				Pass: o.WDB.Pass,
			},
		}
	}
	return o, nil
}

// String 打印可输出的配置
func (s *Source) String() string {
	return fmt.Sprintf("host:%s user:%s", s.Host, s.User)
}

// String 打印可输出的配置
func (c *Config) String() string {
	var str strings.Builder
	fmt.Fprintln(&str, "mysql confiy:")
	fmt.Fprintln(&str, "wdb:", c.WDB)
	fmt.Fprintln(&str, "rdbs:", c.RDBs)
	fmt.Fprintln(&str, "dbname:", c.DBName)
	fmt.Fprintln(&str, "max_open_conn:", c.MaxOpenConn)
	fmt.Fprintln(&str, "max_idle_conn:", c.MaxIdleConn)
	fmt.Fprintln(&str, "max_lifetime:", c.MaxLifetime)
	fmt.Fprintln(&str, "keepalive:", c.Keepalive)
	return str.String()
}
