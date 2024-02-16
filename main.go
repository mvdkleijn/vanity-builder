package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

var config *Config
var routes map[string]Module

type Config struct {
	Outputdirectory string   `yaml:"outputdirectory"`
	DefaultBranch   string   `yaml:"defaultbranch"`
	Vanitydomain    string   `yaml:"vanitydomain"`
	Modules         []Module `yaml:"modules"`
}

type Module struct {
	Name        string   `yaml:"name"`
	Repository  string   `yaml:"repository"`
	Version     string   `yaml:"version"`
	Branch      string   `yaml:"branch"`
	Homepage    string   `yaml:"homepage"`
	Subpackages []string `yaml:"subpackages"`
	// Vanitydomain     string
	Vanitymodulename string
	Versions         map[string]string
}

func (m Module) GetRoutes() (routes []string) {
	routes = append(routes, fmt.Sprintf("/%s", m.Name))

	for _, subpackage := range m.Subpackages {
		routes = append(routes, fmt.Sprintf("/%s/%s", m.Name, subpackage))
	}

	return
}

type TemplateData struct {
	DefaultBranch string
	Vanitydomain  string
	Module        Module
}

func loadConfig() *Config {
	var c Config
	data, err := os.ReadFile("packages.yaml")

	if err != nil {
		log.Fatal().Err(err).Msg("error reading packages.yaml")
	}

	err = yaml.Unmarshal(data, &c)

	if err != nil {
		log.Fatal().Err(err).Msg("error unmarshalling packages.yaml")
	}

	return &c
}

func loadRoutes() (routes map[string]Module) {
	routes = map[string]Module{}

	for _, module := range config.Modules {
		for _, route := range module.GetRoutes() {
			routes[route] = module
		}
	}

	return
}

func main() {
	config = loadConfig()
	routes = loadRoutes()

	t, err := template.ParseFiles("templates/package.tmpl")
	if err != nil {
		log.Err(err).Msg("error parsing package template")
		return
	}

	for route, module := range routes {
		log.Info().Msgf("generating %s", route)

		pkgDir := filepath.Join(config.Outputdirectory, route)

		if err := os.MkdirAll(pkgDir, 0755); err != nil {
			log.Err(err).Msg("error creating route")
			return
		}

		module.Repository = strings.TrimSuffix(module.Repository, ".git")
		module.Vanitymodulename = fmt.Sprintf("%s/%s", config.Vanitydomain, module.Name)

		var page bytes.Buffer
		data := TemplateData{
			DefaultBranch: config.DefaultBranch,
			Vanitydomain:  config.Vanitydomain,
			Module:        module,
		}
		if err := t.Execute(&page, data); err != nil {
			log.Err(err).Msg("error executing template")
			return
		}

		pkgFile := filepath.Join(pkgDir, "index.html")
		if err := os.WriteFile(pkgFile, page.Bytes(), 0755); err != nil {
			log.Err(err).Msg("error writing package file")
			return
		}
	}

	t, err = template.ParseFiles("templates/index.tmpl")
	if err != nil {
		log.Err(err).Msg("error parsing index template")
		return
	}

	var page bytes.Buffer
	if err := t.Execute(&page, config); err != nil {
		log.Err(err).Msg("error executing template")
		return
	}

	pkgFile := filepath.Join(config.Outputdirectory, "index.html")
	if err := os.WriteFile(pkgFile, page.Bytes(), 0755); err != nil {
		log.Err(err).Msg("error writing index file")
		return
	}

}
