#include "greeter.h"
#include <gtest/gtest.h>

using namespace projects::cpp::greeter;

TEST(GreeterTest, DefaultGreet) { EXPECT_EQ(greet(), "Hello World"); }

TEST(GreeterTest, CustomGreet) { EXPECT_EQ(greet("Alice"), "Hello Alice"); }
