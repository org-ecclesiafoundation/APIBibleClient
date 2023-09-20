.PHONY: test

TEST_CMD := go test

CLIENT_PATH := ./pkg/client
PARAMS_PATH := ./pkg/client/params

.DEFAULT:
	$(MAKE) help

help:
	@echo "Usage: TODO"

test:
	$(MAKE) test_client
	$(MAKE) test_params

test_client:
	@echo "Testing Client:"
	$(TEST_CMD) $(CLIENT_PATH)

test_params:
	@echo "Testing Params:"
	$(TEST_CMD) $(PARAMS_PATH)
