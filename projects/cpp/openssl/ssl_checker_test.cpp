#include "projects/cpp/openssl/ssl_checker.h"
#include <gtest/gtest.h>

TEST(SslCheckerTest, VersionStringIsPopulated) {
  EXPECT_FALSE(projects::cpp::openssl::GetLibraryVersion().empty());
}

TEST(SslCheckerTest, IsLinkedToBoringSsl) {
  EXPECT_TRUE(projects::cpp::openssl::IsBoringSsl())
      << "Expected BoringSSL but found: "
      << projects::cpp::openssl::GetLibraryVersion();
}
