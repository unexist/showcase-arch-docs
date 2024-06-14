//
// @package Showcase-Architecture-Documentation
//
// @file Structurizr base
// @copyright 2024-present Christoph Kappel <christoph@unexist.dev>
// @version $Id$
//
// This program can be distributed under the terms of the Apache License v2.0.
// See the file LICENSE for details.
//

package main

import (
	structScaper "github.com/krzysztofreczek/go-structurizr/pkg/scraper"
	structView "github.com/krzysztofreczek/go-structurizr/pkg/view"

	"github.com/unexist/showcase-architecture-documentation/adapter"
	"github.com/unexist/showcase-architecture-documentation/domain"
	"github.com/unexist/showcase-architecture-documentation/infrastructure"
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
