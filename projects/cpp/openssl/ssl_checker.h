#pragma once

#include <string>

namespace projects::cpp::openssl {

// Returns the raw version string from the linked library
auto GetLibraryVersion() -> std::string;

// Returns true if the linked library identifies as BoringSSL
auto IsBoringSsl() -> bool;

} // namespace projects::cpp::openssl
