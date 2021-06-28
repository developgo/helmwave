package plan

import (
	"github.com/helmwave/helmwave/pkg/version"
	log "github.com/sirupsen/logrus"
	"os"
)

func (p *Plan) Import() error {
	body, err := NewBody(p.fullPath)
	if err != nil {
		return err
	}

	err = p.importManifest()
	if err == ErrManifestDirEmpty {
		log.Warn(err)
	}

	if err != ErrManifestDirEmpty && err != nil {
		return err
	}

	p.body = body
	version.Check(p.body.Version, version.Version)

	return nil
}

func (p *Plan) importManifest() error {
	ls, err := os.ReadDir(p.dir + Manifest)
	if err != nil {
		return err
	}

	if len(ls) == 0 {
		return ErrManifestDirEmpty
	}

	for _, l := range ls {
		if !l.IsDir() {
			c, err := os.ReadFile(p.dir + Manifest + l.Name())
			if err != nil {
				return err
			}

			p.manifests[l.Name()] = string(c)
		}
	}

	return nil
}

func (p *Plan) Clean() {
	_ = os.RemoveAll(p.dir)
	_ = os.RemoveAll(p.fullPath)
}
