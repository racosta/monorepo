#include "projects/cpp/openssl/ssl_checker.h"
#include <openssl/crypto.h>   // Include for OpenSSL_version() function
#include <openssl/opensslv.h> // Include OpenSSL version header

namespace projects::cpp::openssl {

std::string GetLibraryVersion() { return OpenSSL_version(OPENSSL_VERSION); }

bool IsBoringSsl() {
  return GetLibraryVersion().find("BoringSSL") != std::string::npos;
}

} // namespace projects::cpp::openssl
