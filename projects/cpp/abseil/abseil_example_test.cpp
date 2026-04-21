#include "projects/cpp/abseil/abseil_example.h"
#include <gtest/gtest.h>

TEST(AbseilExampleTest, JoinItemsReturnsCorrectlyJoinedString) {
  std::vector<std::string> items = {"apple", "banana", "cherry"};
  std::string result = projects::cpp::abseil::JoinItems(items);
  EXPECT_EQ(result, "apple, banana, cherry");
}
