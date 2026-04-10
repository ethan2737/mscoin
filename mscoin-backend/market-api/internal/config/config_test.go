package config

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

func TestConfigLoadsExchangeMysql(t *testing.T) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to resolve current file path")
	}
	configPath := filepath.Join(filepath.Dir(currentFile), "..", "..", "etc", "conf.yaml")
	var c Config
	conf.MustLoad(configPath, &c)
	if c.ExchangeMysql.DataSource == "" {
		t.Fatal("expected exchange mysql datasource to be loaded from conf.yaml")
	}
}
