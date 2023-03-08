package types

type Configuration struct {
	Port     string       `mapstructure:"Port"`
	Links    []Link       `mapstructure:"Links"`
	MetaData PageMetaData `mapstructure:"metaData"`
	Badges   []Badge      `mapstructure:"Badges"`
	Filter   map[string]string
}
