
testCmd = echo "Running test $(1)" \
	&& cat tests/assets/$(2).txt \
		| go run main.go --label $(2) tests/$(1).in.md $(3) > .tmp/out.md \
	&& diff .tmp/out.md tests/$(1).out.md || exit 1

test: clean
	mkdir -p .tmp
	@$(call testCmd,basic,single-line)
	@$(call testCmd,multiple,single-line)
	@$(call testCmd,empty,single-line)
	@$(call testCmd,replace,single-line)
	@$(call testCmd,formatter,single-line)
	@$(call testCmd,multiline,multi-line)
	@$(call testCmd,mismatch,multi-line)

# make test_single test=multiple
test_single:
	$(call testCmd,$(test),single-line)

clean:
	rm -rf dist .tmp

docs:
	go run main.go -h | go run main.go --label help README.md -i
