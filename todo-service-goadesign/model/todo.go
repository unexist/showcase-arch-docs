//
// @package Showcase-Architecture-Documentation
//
// @file Todo model
// @copyright 2024-present Christoph Kappel <christoph@unexist.dev>
// @version $Id$
//
// This program can be distributed under the terms of the Apache License v2.0.
// See the file LICENSE for details.
//

package model

import . "goa.design/model/dsl"

var backend_container interface{}

var _ = Design("Todo design", "This is a design of the todo service.", func() {
	Version("0.1")

	var system = SoftwareSystem("Software System", "The software system.", func() {
		Tag("system")

		Container("Frontend Container", "A frontend container", "Angular", func() {
			Tag("application")
			URL("https://unexist.blog/todo")
			Uses("Backend Container", "Uses", "HTTP", Synchronous)
			Delivers("User", "Handles todo from", "HTTP", Asynchronous)
		})

		backend_container = Container("Backend Container", "A backend container", "GinTonic + REST", func() {
			Tag("application")
			URL("https://unexist.blog/todo")
			Uses("Database Container", "Persists and retrieves data", "SQL", Synchronous)
			Delivers("User", "Handles todo from", "HTTP", Asynchronous)
			Component("Todo Service", "Golang")
		})

		Container("Database Container", "A database container", "Postgresql", func() {
			Tag("infrastructure")
			URL("https://unexist.blog/todo")
		})
	})

	DeploymentEnvironment("Dev", func() {
		DeploymentNode("Cloud", func() {
			ContainerInstance("Software System/Frontend Container")
			ContainerInstance("Software System/Backend Container")
			ContainerInstance("Software System/Database Container")
		})
	})

	Person("User", "A user of the software system.", func() {
		Uses(system, "Uses")
		Tag("person")
	})

	Person("Admin", "A user of the software system.", func() {
		Uses(system, "Maintains")
		Tag("person")
	})

	Views(func() {
		SystemLandscapeView("SystemLandscapeView", "A System Landscape View", func() {
			Title("Overview of the system landscape")
			AddAll()
			AutoLayout(RankLeftRight)
		})

		SystemContextView(system, "SystemContext", "A System Context diagram.", func() {
			Title("Overview of the system")
			AddAll()
			AutoLayout(RankLeftRight)
		})

		ContainerView(system, "ContainerView", "A Container View", func() {
			Title("Overview of the containers")
			AddAll()
			AutoLayout(RankLeftRight)
		})

		ComponentView(backend_container, "ComponentView", "A Component View", func() {
			Title("Overview of the components")
			AddAll()
			AutoLayout(RankLeftRight)
		})

		DeploymentView(Global, "Dev", "deployment", "A Deployment View", func() {
			Title("Overview of the deployment on Dev")
			AddAll()
			AutoLayout(RankLeftRight)
		})

		Styles(func() {
			ElementStyle("system", func() {
				Background("#1168bd")
				Color("#ffffff")
			})
			ElementStyle("person", func() {
				Shape(ShapePerson)
				Background("#08427b")
				Color("#ffffff")
			})
		})
	})
})
