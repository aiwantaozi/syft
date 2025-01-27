package config

import (
	"github.com/anchore/syft/syft/pkg/cataloger"
	"github.com/spf13/viper"
)

type pkg struct {
	Cataloger               catalogerOptions  `yaml:"cataloger" json:"cataloger" mapstructure:"cataloger"`
	SearchByBuildTools      buildtoolsOptions `yaml:"search-by-buildtools" json:"search-by-buildtools" mapstructure:"search-by-buildtools"`
	SearchUnindexedArchives bool              `yaml:"search-unindexed-archives" json:"search-unindexed-archives" mapstructure:"search-unindexed-archives"`
	SearchIndexedArchives   bool              `yaml:"search-indexed-archives" json:"search-indexed-archives" mapstructure:"search-indexed-archives"`
}

func (cfg pkg) loadDefaultValues(v *viper.Viper) {
	cfg.Cataloger.loadDefaultValues(v)
	cfg.SearchByBuildTools.loadDefaultValues(v)
	c := cataloger.DefaultSearchConfig()
	v.SetDefault("package.search-unindexed-archives", c.IncludeUnindexedArchives)
	v.SetDefault("package.search-indexed-archives", c.IncludeIndexedArchives)
}

func (cfg *pkg) parseConfigValues() error {
	return cfg.Cataloger.parseConfigValues()
}

func (cfg pkg) ToConfig() cataloger.Config {
	return cataloger.Config{
		Search: cataloger.SearchConfig{
			ByBuildTools:             cfg.SearchByBuildTools.Enabled,
			ByBuildToolsWithMode:     cfg.SearchByBuildTools.Mode,
			IncludeIndexedArchives:   cfg.SearchIndexedArchives,
			IncludeUnindexedArchives: cfg.SearchUnindexedArchives,
			Scope:                    cfg.Cataloger.ScopeOpt,
		},
	}
}
