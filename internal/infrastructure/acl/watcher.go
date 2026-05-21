package acl

import (
	"log"
	"path/filepath"

	"github.com/Citrus0974/ProxyProject/internal/usecase/acl"
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	watcher *fsnotify.Watcher

	loader  *Loader
	service *acl.Service

	defaultAllow bool
}

func NewWatcher(loader *Loader, service *acl.Service, defaultAllow bool) (*Watcher, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Watcher{
		watcher:      w,
		loader:       loader,
		service:      service,
		defaultAllow: defaultAllow,
	}, nil
}

func (w *Watcher) Start(path string) error {

	dir := filepath.Dir(path)

	err := w.watcher.Add(dir)
	if err != nil {
		return err
	}

	go func() {

		for {

			select {

			case event := <-w.watcher.Events:

				if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Rename) == 0 {
					continue
				}

				rules, err := w.loader.Load(
					w.defaultAllow,
				)

				if err != nil {
					log.Println(err)
					continue
				}

				w.service.ReplaceRules(rules)

			case err := <-w.watcher.Errors:
				log.Println(err)
			}
		}
	}()

	return nil
}
