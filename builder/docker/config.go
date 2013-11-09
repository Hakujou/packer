package docker

import (
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/packer"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	ExportPath string `mapstructure:"export_path"`
	Image      string

	tpl *packer.ConfigTemplate
}
