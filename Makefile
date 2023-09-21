COMPOSEFLAGS=-d
ITESTS_L2_HOST=http://localhost:9545
BEDROCK_TAGS_REMOTE?=origin
# Requires at least Python v3.9; specify a minor version below if needed
PYTHON=python3

build: build-go build-ts
.PHONY: build

build-go: submodules op-node op-proposer op-batcher op-erigon
.PHONY: build-go

build-ts: submodules
	if [ -n "$$NVM_DIR" ]; then \
		. $$NVM_DIR/nvm.sh && nvm use; \
	fi
	pnpm install
	pnpm build
.PHONY: build-ts

ci-builder:
	docker build -t ci-builder -f ops/docker/ci-builder/Dockerfile .

submodules:
	# CI will checkout submodules on its own (and fails on these commands)
	if [ -z "$$GITHUB_ENV" ]; then \
		echo "SKIP submodules" \
		#git submodule init; \
		#git submodule update; \
	fi
.PHONY: submodules

op-bindings:
	make -C ./op-bindings
.PHONY: op-bindings

op-node:
	make -C ./op-node op-node
.PHONY: op-node

generate-mocks-op-node:
	make -C ./op-node generate-mocks
.PHONY: generate-mocks-op-node

generate-mocks-op-service:
	make -C ./op-service generate-mocks
.PHONY: generate-mocks-op-service

op-batcher:
	make -C ./op-batcher op-batcher
.PHONY: op-batcher

op-proposer:
	make -C ./op-proposer op-proposer
.PHONY: op-proposer

op-challenger:
	make -C ./op-challenger op-challenger
.PHONY: op-challenger

op-program:
	make -C ./op-program op-program
.PHONY: op-program

op-erigon:
	make -C ./op-erigon erigon
.PHONY: op-erigon

cannon:
	make -C ./cannon cannon
.PHONY: cannon

cannon-prestate: op-program cannon
	./cannon/bin/cannon load-elf --path op-program/bin/op-program-client.elf --out op-program/bin/prestate.json --meta op-program/bin/meta.json
	./cannon/bin/cannon run --proof-at '=0' --stop-at '=1' --input op-program/bin/prestate.json --meta op-program/bin/meta.json --proof-fmt 'op-program/bin/%d.json' --output ""
	mv op-program/bin/0.json op-program/bin/prestate-proof.json

mod-tidy:
	# Below GOPRIVATE line allows mod-tidy to be run immediately after
	# releasing new versions. This bypasses the Go modules proxy, which
	# can take a while to index new versions.
	#
	# See https://proxy.golang.org/ for more info.
	export GOPRIVATE="github.com/ethereum-optimism,github.com/bobanetwork"
	export GOPROXY="https://proxy.golang.org,direct"
	go mod tidy
.PHONY: mod-tidy

clean:
	rm -rf ./bin
	make -C ./op-erigon clean
.PHONY: clean

nuke: clean devnet-clean
	git clean -Xdf
.PHONY: nuke

op-program/bin/prestate-proof.json:
	make cannon-prestate

devnet-up: op-program/bin/prestate-proof.json
	$(shell ./ops/scripts/newer-file.sh .devnet/allocs-l1.json ./packages/contracts-bedrock)
	if [ $(.SHELLSTATUS) -ne 0 ]; then \
		make devnet-allocs; \
	fi
	PYTHONPATH=./bedrock-devnet ${PYTHON} ./bedrock-devnet/main.py --monorepo-dir=.
.PHONY: devnet-up

# alias for devnet-up
devnet-up-deploy: devnet-up

devnet-hardhat-up: op-program/bin/prestate-proof.json
	PYTHONPATH=./bedrock-devnet ${PYTHON} ./bedrock-devnet/hardhat.py --monorepo-dir=.
.PHONY: devnet-hardhat-up

devnet-hardhat-test:
	PYTHONPATH=./bedrock-devnet ${PYTHON} ./bedrock-devnet/hardhat.py --monorepo-dir=. --test
.PHONY: devnet-hardhat-test

devnet-test:
	PYTHONPATH=./bedrock-devnet ${PYTHON} ./bedrock-devnet/main.py --monorepo-dir=. --test
.PHONY: devnet-test

devnet-down:
	@(cd ./ops-bedrock && GENESIS_TIMESTAMP=$(shell date +%s) docker compose stop)
.PHONY: devnet-down

devnet-clean:
	rm -rf ./packages/contracts-bedrock/deployments/devnetL1
	rm -rf ./.devnet
	cd ./ops-bedrock && docker compose down
	docker image ls 'ops-bedrock*' --format='{{.Repository}}' | xargs -r docker rmi
	docker volume ls --filter name=ops-bedrock --format='{{.Name}}' | xargs -r docker volume rm
.PHONY: devnet-clean

devnet-allocs:
	PYTHONPATH=./bedrock-devnet ${PYTHON} ./bedrock-devnet/main.py --monorepo-dir=. --allocs

devnet-logs:
	@(cd ./ops-bedrock && docker compose logs -f)
	.PHONY: devnet-logs

test-unit:
	make -C ./op-node test
	make -C ./op-proposer test
	make -C ./op-batcher test
	make -C ./op-e2e test
	pnpm test
.PHONY: test-unit

test-integration:
	bash ./ops-bedrock/test-integration.sh \
		./packages/contracts-bedrock/deployments/devnetL1
.PHONY: test-integration

# Remove the baseline-commit to generate a base reading & show all issues
semgrep:
	$(eval DEV_REF := $(shell git rev-parse develop))
	SEMGREP_REPO_NAME=ethereum-optimism/optimism semgrep ci --baseline-commit=$(DEV_REF)
.PHONY: semgrep

clean-node-modules:
	rm -rf node_modules
	rm -rf packages/**/node_modules

tag-bedrock-go-modules:
	./ops/scripts/tag-bedrock-go-modules.sh $(BEDROCK_TAGS_REMOTE) $(VERSION)
.PHONY: tag-bedrock-go-modules

update-op-geth:
	./ops/scripts/update-op-geth.py
.PHONY: update-op-geth

bedrock-markdown-links:
	docker run --init -it -v `pwd`:/input lycheeverse/lychee --verbose --no-progress --exclude-loopback \
		--exclude twitter.com --exclude explorer.optimism.io --exclude linux-mips.org \
		--exclude-mail /input/README.md "/input/specs/**/*.md"

install-geth:
	go install github.com/ethereum/go-ethereum/cmd/geth@v1.12.0
