# Show the help with the list of commands
help:
  @just --list

# Clean Bazel output files
clean:
  bazel clean

# Clean and purge all Bazel files
pristine:
  bazel clean --expunge

TEST_OUTPUT_OPTIONS := "--test_output=errors"

# Run Bazel test
test:
  bazel test {{TEST_OUTPUT_OPTIONS}} ...
