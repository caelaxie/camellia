package camellia

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	internalcamellia "github.com/caelaxie/camellia/internal/camellia"
)

func init() {
	register.Plugin("camellia", New)
}

type settings struct{}

type linterPlugin struct{}

func New(rawSettings any) (register.LinterPlugin, error) {
	if rawSettings != nil {
		if _, err := register.DecodeSettings[settings](rawSettings); err != nil {
			return nil, err
		}
	}

	return &linterPlugin{}, nil
}

func (p *linterPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{internalcamellia.Analyzer}, nil
}

func (p *linterPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
