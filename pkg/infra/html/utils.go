package html

import (
	"fmt"
	"strings"
)

type Web struct {
	Styles     string `yaml:"styles"` // actual css
	HTML       string `yaml:"html"`   // actual html
	Components map[string]string
	Route      string `yaml:"route"`
}

func (wc *Web) Compile() string {
	ht := strings.ReplaceAll(wc.HTML, `<style></style>`, fmt.Sprintf(`<style>%s</style>`, wc.Styles))

	// for each component
	for key, val := range wc.Components {
		ht = strings.ReplaceAll(wc.HTML, key, val)
	}

	return ht
}

type HTMLProvider struct {
	WebComponents map[string]*Web `yaml:"web_components"`
}

func InitialiseProvider() (*HTMLProvider, error) {

	ht := &HTMLProvider{
		WebComponents: make(map[string]*Web),
	}

	return ht, nil
}

func (g *HTMLProvider) AddComponent(route string, style string, html string, comp map[string]string) error {

	g.WebComponents[route] = &Web{
		Styles:     style,
		HTML:       html,
		Route:      route,
		Components: comp,
	}

	return nil
}
