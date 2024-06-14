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

var _ = Design("Getting Started", "This is a model of my software system.", func() {
	var System = SoftwareSystem("Software System", "My software system.", func() {
		Tag("system")
	})

	Person("User", "A user of my software system.", func() {
		Uses(System, "Uses")
		Tag("person")
	})

	Views(func() {
		SystemContextView(System, "SystemContext", "An example of a System Context diagram.", func() {
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
