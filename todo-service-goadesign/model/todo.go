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

var (
	container_frontend interface{}
	container_backend  interface{}
)

var _ = Design("Todo design", "This is a design of the todo service", func() {
	Version("0.1")
	Enterprise("Todo Showcase Service")

	var system = SoftwareSystem("Software System", "The todo software system", func() {
		Tag("system")
		URL("https://unexist.blog")

		Container("Webserver", "A webserver to deliver the frontend", "Nginx", func() {
			Tag("infrastructure")
			URL("https://nginx.org/")

			Delivers("User", "Handles requests from", "HTTP", Asynchronous)
			Delivers("User", "Delivers frontend to", "HTTP", Asynchronous)
		})

		container_frontend = Container("Web Frontend", "A Angular-based web frontend", "Angular + REST", func() {
			Tag("frontend")

			Uses("Webserver", "Is delivered by", "HTTP", Asynchronous)
			Uses("Web API", "Makes API calls to", "HTTP", Asynchronous)
		})

		container_backend = Container("Web API", "A backend service", "GinTonic + REST", func() {
			Tag("backend")

			Uses("Database", "Reads from and writes to", "SQL/TCP", Asynchronous)

			Component("Todo Service", "Domain logic for todo", "Golang", func() {
				Tag("service")
			})
		})

		Container("Database", "A RDBMS to handle the data", "Postgresql", func() {
			Tag("infrastructure")
			URL("https://postgresql.org")
		})
	})

	DeploymentEnvironment("Dev", func() {
		DeploymentNode("Cloud", func() {
			ContainerInstance("Software System/Webserver")
			ContainerInstance("Software System/Web Frontend")
			ContainerInstance("Software System/Web API")
			ContainerInstance("Software System/Database")
		})
	})

	Person("User", "A user of the software system.", func() {
		Tag("person")

		Uses(system, "Uses")
	})

	Views(func() {
		SystemLandscapeView("SystemLandscapeView", "A System Landscape View", func() {
			Title("Overview of the system landscape")
			AddAll()
		})

		SystemContextView(system, "SystemContext", "A System Context diagram.", func() {
			Title("Overview of the system")
			AddAll()
		})

		ContainerView(system, "ContainerView", "A Container View", func() {
			Title("Overview of the containers")
			AddAll()
		})

		ComponentView(container_frontend, "ComponentView Frontend", "A Component View of the web frontend", func() {
			Title("Overview of the frontend components")
			AddComponents()
		})

		ComponentView(container_backend, "ComponentView Backend", "A Component View of the web backend", func() {
			Title("Overview of the backend components")
			AddComponents()
		})

		DeploymentView(Global, "Dev", "deployment", "A Deployment View", func() {
			Title("Overview of the deployment on Dev")
			AddAll()
		})

		Styles(func() {
			ElementStyle("system", func() {
				Background("#1168bd")
				Color("#ffffff")
			})

			ElementStyle("frontend", func() {
				Shape(ShapeWebBrowser)
			})

			ElementStyle("backend", func() {
				Shape(ShapeRoundedBox)
			})

			ElementStyle("service", func() {
				Shape(ShapeHexagon)
			})

			ElementStyle("infrastructure", func() {
				Shape(ShapeComponent)
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
