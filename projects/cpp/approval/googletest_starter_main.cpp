#define APPROVALS_GOOGLETEST
#include "ApprovalTests.hpp"
#include <gtest/gtest.h>

class ApprovalTestsSetup : public ::testing::Environment {
public:
  ~ApprovalTestsSetup() override {}

  void SetUp() override {
    ApprovalTests::ApprovalTestNamer::setCheckBuildConfig(false);

    static auto directoryDisposer =
        ApprovalTests::Approvals::useApprovalsSubdirectory("approval_tests");

    // This will print the command to diff the files to the console if the test
    // fails.
    static auto reporterDisposer =
        ApprovalTests::Approvals::useAsDefaultReporter(
            std::make_shared<ApprovalTests::TextDiffReporter>());
  }
};

auto *approval_env =
    ::testing::AddGlobalTestEnvironment(new ApprovalTestsSetup);
