#define APPROVALS_GOOGLETEST
#include "ApprovalTests.hpp"
#include <gtest/gtest.h>

// This class handles the setup for all tests in the binary
class ApprovalTestsSetup : public ::testing::Environment {
public:
  ~ApprovalTestsSetup() override {}

  void SetUp() override {
    // 1. Keep the path fix from before
    ApprovalTests::ApprovalTestNamer::setCheckBuildConfig(false);

    // 2. Set the subdirectory
    static auto directoryDisposer =
        ApprovalTests::Approvals::useApprovalsSubdirectory("approval_tests");

    // 3. Configure a Reporter for Bazel
    // This will print the command to diff the files to the console
    // if the test fails.
    static auto reporterDisposer =
        ApprovalTests::Approvals::useAsDefaultReporter(
            std::make_shared<ApprovalTests::TextDiffReporter>());
  }
};

// Register the environment
// This executes before any TEST() starts
auto *approval_env =
    ::testing::AddGlobalTestEnvironment(new ApprovalTestsSetup);
