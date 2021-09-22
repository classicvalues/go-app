// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	f, err := os.Create("scripts.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "//go:build !wasm")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "package app")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "// Code generated by go generate; DO NOT EDIT.")
	fmt.Fprintln(f)

	gen := []struct {
		Var      string
		Filename string
	}{
		{Var: "wasmExecJS", Filename: filepath.Join(
			runtime.GOROOT(),
			"misc",
			"wasm",
			"wasm_exec.js",
		)},
		{Var: "appJS", Filename: "gen/app.js"},
		{Var: "appWorkerJS", Filename: "gen/app-worker.js"},
		{Var: "manifestJSON", Filename: "gen/manifest.webmanifest"},
		{Var: "appCSS", Filename: "gen/app.css"},
	}

	fmt.Fprintln(f, "const(")

	for _, g := range gen {
		b, err := ioutil.ReadFile(g.Filename)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(f, "%s = %q", g.Var, b)
		fmt.Fprintln(f)
		fmt.Fprintln(f)
	}

	fmt.Fprintln(f, ")")
}
