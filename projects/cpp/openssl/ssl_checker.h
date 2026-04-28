#pragma once

#include <string>

namespace projects::cpp::openssl {

// Returns the raw version string from the linked library
std::string GetLibraryVersion();

// Returns true if the linked library identifies as BoringSSL
bool IsBoringSsl();

} // namespace projects::cpp::openssl
