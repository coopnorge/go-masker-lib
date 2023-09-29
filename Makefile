# SPDX-FileCopyrightText: 2022 Coop Norge SA
#
# SPDX-License-Identifier: MIT

# vim: set noexpandtab fo-=t:
# https://www.gnu.org/software/make/manual/make.html
SHELL=bash
.SHELLFLAGS=-ec -o pipefail
current_makefile:=$(lastword $(MAKEFILE_LIST))
current_makefile_dir:=$(dir $(abspath $(current_makefile)))

.PHONY: all
all:

########################################################################
# boiler plate
########################################################################
SHELL=bash

ifneq ($(filter all vars,$(VERBOSE)),)
dump_var=$(info var $(1)=$($(1)))
dump_vars=$(foreach var,$(1),$(call dump_var,$(var)))
else
dump_var=
dump_vars=
endif

ifneq ($(filter all targets,$(VERBOSE)),)
__ORIGINAL_SHELL:=$(SHELL)
SHELL=$(warning Building $@$(if $<, (from $<))$(if $?, ($? newer)))$(TIME) $(__ORIGINAL_SHELL)
endif

define __newline


endef

########################################################################
# variables ...
########################################################################
GOPATH:=$(shell type go >/dev/null 2>&1 && go env GOPATH)
export PATH:=$(if $(GOPATH),$(GOPATH)/bin:,)$(PATH)
export GOFLAGS:=-mod=vendor

PROJECT_TOOLS_DIR?=$(current_makefile_dir)var/tools/
tools_dir=$(PROJECT_TOOLS_DIR)
ifneq ($(tools_dir),)
export PATH:=$(tools_dir):$(PATH)
endif
XDG_CACHE_HOME?=$(HOME)/.cache

########################################################################
# targets ...
########################################################################

.PHONY: clean
clean: ## clean build outputs
clean:
	rm -vr coverage.out || :
	go clean

.PHONY: distclean
distclean: clean-$(gopath_local)/ ## restore repo to pristine state

.PHONY: validate-static
validate-static: ## run static validation
	$(golangci_lint_bin) run $(if $(filter all commands,$(VERBOSE)),-v) ./...

.PHONY: validate-fix
validate-fix: ## fix auto-fixable validation errors
	$(golangci_lint_bin) run $(if $(filter all commands,$(VERBOSE)),-v) --fix ./...

.PHONY: test
test:  ## run tests
test coverage.out:
	go test -cover -race \
		-coverprofile=coverage.out -covermode=atomic \
		$(if $(filter all commands,$(VERBOSE)),-v) \
		$(if $(gotest_files),$(gotest_files),./...) \
		$(gotest_args)

.PHONY: view-coverage
view-coverage: coverage.out ## view coverage
	go tool cover -html=$(<)

.PHONY: validate
validate: validate-static test ## validate everything

.PHONY: watch
watch: $(modd_bin)
	$(modd_bin) --debug --notify --file modd.conf

.PHONY: publish-coverage
publish-coverage: coverage.out ## publish code coverage
	$(codecov_bin)

########################################################################
# toolchain
########################################################################

golangci_lint_bin=$(tools_dir)golangci-lint
golangci_lint_version=1.54.2
golangci_lint_qualified=golangci-lint-$(golangci_lint_version)-linux-amd64.tar.gz
$(golangci_lint_bin): | $(tools_dir) $(XDG_CACHE_HOME)/
	cd $(XDG_CACHE_HOME) \
		&& wget -c -O $(golangci_lint_qualified) \
			https://github.com/golangci/golangci-lint/releases/download/v$(golangci_lint_version)/$(golangci_lint_qualified) \
		&& sha256sum -c $(abspath $(golangci_lint_qualified).sha256sum)
	tar -zxvf $(XDG_CACHE_HOME)/$(golangci_lint_qualified) --strip-components=1 -C $(tools_dir)

modd_bin=$(tools_dir)modd
modd_version=0.8
modd_qualified=modd-$(modd_version)-linux64.tgz
$(modd_bin): | $(tools_dir)/ $(XDG_CACHE_HOME)/
	cd $(XDG_CACHE_HOME) \
		&& wget -c -O $(modd_qualified) \
			https://github.com/cortesi/modd/releases/download/v$(modd_version)/$(modd_qualified) \
		&& sha256sum -c $(abspath $(modd_qualified).sha256sum)
	tar -zxvf $(XDG_CACHE_HOME)/$(modd_qualified) --strip-components=1 -C $(tools_dir)

codecov_bin=$(tools_dir)codecov
codecov_version=0.1.9
codecov_qualified=codecov-$(codecov_version)
$(codecov_bin): | $(tools_dir)/ $(XDG_CACHE_HOME)/
	cd $(XDG_CACHE_HOME) \
		&& wget -c -O $(codecov_qualified) \
			https://uploader.codecov.io/v$(codecov_version)/linux/codecov \
		&& sha256sum -c $(abspath $(codecov_qualified).sha256sum) \
		&& chmod +x $(codecov_qualified)
	cp $(XDG_CACHE_HOME)/$(codecov_qualified) $(codecov_bin)

.PHONY: toolchain
toolchain: ## install toolchain
toolchain: $(golangci_lint_bin) $(modd_bin) $(codecov_bin)

.PHONY: distclean
distclean: ## restore repo to pristine state
distclean: clean-$(tools_dir)/

########################################################################
# useful ...
########################################################################
## force ...

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: .FORCE
.FORCE:
$(force_targets): .FORCE

## dirs ...
.PRECIOUS: %/
%/:
	mkdir -vp $(@)

.PHONY: clean-%/
clean-%/:
	@{ test -d $(*) && { set -x; rm -vr $(*); set +x; } } || echo "directory $(*) does not exist ... nothing to clean"
