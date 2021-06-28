package release

import (
	"github.com/helmwave/helmwave/pkg/template"
	log "github.com/sirupsen/logrus"
	"os"
)

func (rel *Config) RenderValues(dir string) error {
	rel.filterValuesFiles()

	for i, v := range rel.Values {

		s := v + "." + rel.UniqName() + ".plan"

		p := dir + s
		err := template.Tpl2yml(v, p, struct{ Release *Config }{rel})
		if err != nil {
			return err
		}

		rel.Values[i] = p
	}

	return nil
}

// TODO: add context
// filterValuesFiles filters non-existent values files.
func (rel *Config) filterValuesFiles() {
	for i := len(rel.Values) - 1; i >= 0; i-- {
		stat, err := os.Stat(rel.Values[i])
		if os.IsNotExist(err) || stat.IsDir() {
			log.Warn(rel.Values[i], " skipping")
			rel.Values = append(rel.Values[:i], rel.Values[i+1:]...)
		}
	}
}
