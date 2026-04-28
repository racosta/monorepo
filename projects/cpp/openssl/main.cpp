#include "projects/cpp/openssl/ssl_checker.h"
#include <spdlog/cfg/argv.h>
#include <spdlog/cfg/env.h>
#include <spdlog/sinks/stdout_color_sinks.h>
#include <spdlog/spdlog.h>

int main(int argc, char *argv[]) {
  // SPDLOG_LEVEL=error
  spdlog::cfg::load_env_levels();
  // ./rapidjson SPDLOG_LEVEL=error
  spdlog::cfg::load_argv_levels(argc, argv);

  spdlog::info("Checking SSL Linkage...");
  spdlog::info("Library Identity: {}",
               projects::cpp::openssl::GetLibraryVersion());

  if (projects::cpp::openssl::IsBoringSsl()) {
    spdlog::info("RESULT: Success. Using BoringSSL.");
    return 0;
  } else {
    spdlog::error("RESULT: Failure. Not linked to BoringSSL.");
    return 1;
  }
}
