#include "spdlog/cfg/argv.h"
#include "spdlog/cfg/env.h"
#include "spdlog/sinks/stdout_color_sinks.h"
#include "spdlog/spdlog.h"
#include <openssl/crypto.h>   // Include for OpenSSL_version() function
#include <openssl/opensslv.h> // Include OpenSSL version header

int main(int argc, char *argv[]) {
  // SPDLOG_LEVEL=error
  spdlog::cfg::load_env_levels();
  // ./rapidjson SPDLOG_LEVEL=error
  spdlog::cfg::load_argv_levels(argc, argv);

  spdlog::info("OpenSSL Version: {}", OpenSSL_version(OPENSSL_VERSION));
  return 0;
}
