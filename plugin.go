package camellia

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	internalcamellia "github.com/caelaxie/camellia/internal"
)

func init() {
	register.Plugin("camellia", New)
}

type settings struct {
	Exclude []string `json:"exclude"`
}

type linterPlugin struct {
	analyzer *analysis.Analyzer
}

func New(rawSettings any) (register.LinterPlugin, error) {
	var decoded settings

	if rawSettings != nil {
		var err error
		decoded, err = register.DecodeSettings[settings](rawSettings)
		if err != nil {
			return nil, err
		}
	}

	analyzer, err := internalcamellia.NewAnalyzer(internalcamellia.Config{
		Exclude: decoded.Exclude,
	})
	if err != nil {
		return nil, err
	}

	return &linterPlugin{analyzer: analyzer}, nil
}

func (p *linterPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{p.analyzer}, nil
}

func (p *linterPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
