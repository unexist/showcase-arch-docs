PKG := github.com/unexist/showcase-architecture-documentation
OUT_FILE := design.json

--guard-%:
	@if [ "${${*}}" = "" ]; then \
		echo "Environment variable $* not set: $*=abc123 make $(MAKECMDGOALS)"; \
		exit 1; \
	fi

serve:
	mdl serve $(PKG)/model -dir gen

generate:
	mdl gen $(PKG)/model -out $(OUT_FILE)

struct-gen:
	stz gen $(PKG)/model -out $(OUT_FILE)

struct-put: --guard-ID --guard-KEY --guard-SECRET struct-gen
	stz put $(OUT_FILE) -id $(ID) -key $(KEY) -secret $(SECRET)

install:
	go install goa.design/model/cmd/mdl@latest
	go install goa.design/model/cmd/stz@latest
