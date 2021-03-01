help:
	@echo
	@echo "***********************************************************"
	@echo "*                                                         *"
	@echo "*  Generate documentation locally:                        *"
	@echo "*                                                         *"
	@echo "*  make docs                                              *"
	@echo "*                                                         *"
	@echo "**********************************************************"
	@echo

# Docs
PDFOUT := ARC42.pdf
HTMLOUT := arc42-template.html

pdf:
	asciidoctor-pdf src/main/asciidoc/arc42-template.adoc -D . -o $(PDFOUT) -a imagesdir=images

	@echo
	@echo "******************************************************"
	@echo "*                                                    *"
	@echo "* Documentation can be found here:                   *"
	@echo "* $(PDFOUT)                                          *"
	@echo "*                                                    *"
	@echo "******************************************************"
	@echo

open-pdf:
	open $(PDFOUT)

open:
	open ./target/static/documentation/$(HTMLOUT)

.DEFAULT_GOAL := docs
docs:
	@mvn -f pom.xml clean generate-resources

	@echo
	@echo "**********************************************************"
	@echo "*                                                        *"
	@echo "* Documentation can be found here:                       *"
	@echo "* ./target/static/documentation/$(HTMLOUT)               *"
	@echo "*                                                        *"
	@echo "**********************************************************"
	@echo

