.PHONY: test

TEST_CMD := go test

CLIENT_PATH := ./pkg/client
PARAMS_PATH := ./pkg/client/params
UTILS_PATH  := ./pkg/utils

.DEFAULT:
	$(MAKE) help

help:
	@echo "Usage: TODO"

test:
	$(MAKE) test_client
	$(MAKE) test_params
	$(MAKE) test_utils

test_client:
	@echo "Testing Client:"
	$(TEST_CMD) $(CLIENT_PATH)

test_params:
	@echo "Testing Params:"
	$(TEST_CMD) $(PARAMS_PATH)

test_utils:
	@echo "Testing Utils:"
	$(TEST_CMD) $(UTILS_PATH)