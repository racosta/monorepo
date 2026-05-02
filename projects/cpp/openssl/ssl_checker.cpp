#include "projects/cpp/openssl/ssl_checker.h"
#include <openssl/crypto.h>   // Include for OpenSSL_version() function
#include <openssl/opensslv.h> // Include OpenSSL version header

namespace projects::cpp::openssl {

auto GetLibraryVersion() -> std::string {
  return OpenSSL_version(OPENSSL_VERSION);
}

auto IsBoringSsl() -> bool {
  return GetLibraryVersion().find("BoringSSL") != std::string::npos;
}

} // namespace projects::cpp::openssl
