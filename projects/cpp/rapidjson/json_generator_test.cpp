#include "projects/cpp/rapidjson/json_generator.h"
#include <gtest/gtest.h>
#include <rapidjson/document.h>

TEST(JsonGeneratorTest, ReturnsValidJsonWithCorrectKeys) {
  std::string output = projects::cpp::rapidjson::GenerateProjectInfo();

  rapidjson::Document d;
  // Verify it's even valid JSON
  ASSERT_FALSE(d.Parse(output.c_str()).HasParseError());

  // Verify specific content
  ASSERT_TRUE(d.HasMember("project"));
  EXPECT_STREQ(d["project"].GetString(), "Bazel Example");

  ASSERT_TRUE(d.HasMember("status"));
  EXPECT_STREQ(d["status"].GetString(), "working");
}
