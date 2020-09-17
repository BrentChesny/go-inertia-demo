package main

import (
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"io/ioutil"

	"github.com/brentchesny/go-inertia-demo/app"
	"github.com/brentchesny/go-inertia/inertia"
)

func main() {

	// Load root template
	rootTmpl := template.Must(template.New("base.tmpl").ParseFiles("resources/templates/layouts/base.tmpl"))

	// Initialize Inertia
	inertia.Init(rootTmpl)
	inertia.VersionFunc(func() string {
		manifest, _ := ioutil.ReadFile("public/mix-manifest.json")
		sum := md5.Sum(manifest)
		return hex.EncodeToString(sum[:])
	})

	// Start our app
	app := app.New()
	app.Run()
}
