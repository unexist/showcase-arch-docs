//
// @package Showcase-Structurizr-Golang
//
// @file Structurizr base
// @copyright 2023-present Christoph Kappel <christoph@unexist.dev>
// @version $Id$
//
// This program can be distributed under the terms of the Apache License v2.0.
// See the file LICENSE for details.
//

package main

import (
	structScaper "github.com/krzysztofreczek/go-structurizr/pkg/scraper"
	structView "github.com/krzysztofreczek/go-structurizr/pkg/view"

	"github.com/unexist/showcase-structurizr-golang/adapter"
	"github.com/unexist/showcase-structurizr-golang/domain"
	"github.com/unexist/showcase-structurizr-golang/infrastructure"
	"os"
)

func main() {
	scraper, err := structScaper.NewScraperFromConfigFile("scraper.yaml")
	if err != nil {
		panic(err)
	}

	/* Create business stuff */
	var todoRepository *infrastructure.TodoFakeRepository

	todoRepository = infrastructure.NewTodoFakeRepository()

	defer todoRepository.Close()

	todoService := domain.NewTodoService(todoRepository)
	todoResource := adapter.NewTodoResource(todoService)

	structure := scraper.Scrape(todoResource)

	view, err := structView.NewViewFromConfigFile("view.yaml")
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create("c4.plantuml")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = view.RenderStructureTo(structure, outFile)
	if err != nil {
		panic(err)
	}
}
