package release

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"helm.sh/helm/v3/pkg/action"
	helm "helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	protobufChart "k8s.io/helm/pkg/proto/hapi/chart"
	protobufRelease "k8s.io/helm/pkg/proto/hapi/release"
	"os"
)

func (rel *Config) Sync() (*release.Release, error) {
	var err error
	rel.cfg, err = rel.newCfg()
	if err != nil {
		return nil, err
	}

	helmClient, err := rel.helm()
	if err != nil {
		return nil, err
	}

	return rel.upgrade(helmClient)
}

func _newProtobufRelease(r *release.Release) *protobufRelease.Release {
	return &protobufRelease.Release{
		Name:     r.Name,
		Info:     nil,
		Chart:    &protobufChart.Chart{},
		Config:   nil,
		Manifest: r.Manifest,
		Hooks:    nil,
		Version:  int32(r.Version),
	}
}

func (rel *Config) newCfg() (*action.Configuration, error) {
	cfg := new(action.Configuration)
	helmDriver := os.Getenv("HELM_DRIVER")
	err := cfg.Init(genericclioptions.NewConfigFlags(false), rel.Namespace, helmDriver, log.Debugf)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (rel *Config) helm() (*helm.EnvSettings, error) {
	env := helm.New()
	fs := &pflag.FlagSet{}
	env.AddFlags(fs)
	flag := fs.Lookup("namespace")
	err := flag.Value.Set(rel.Namespace)
	if err != nil {
		return nil, err
	}

	return env, nil
}
