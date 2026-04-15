#include "projects/cpp/nlohmann_json/json_generator.h"
#include <gtest/gtest.h>
#include <nlohmann/json.hpp>

TEST(JsonGeneratorTest, ReturnsValidJsonWithCorrectKeys) {
  std::string output = projects::cpp::nlohmann_json::GenerateProjectInfo();

  nlohmann::json j;
  // Verify it's even valid JSON
  ASSERT_NO_THROW(j = nlohmann::json::parse(output));

  // Verify specific content
  ASSERT_TRUE(j.contains("project"));
  EXPECT_EQ(j["project"], "Bazel Example");

  ASSERT_TRUE(j.contains("status"));
  EXPECT_EQ(j["status"], "working");
}
