#include "nlohmann/json.hpp"
#include "spdlog/cfg/argv.h"
#include "spdlog/cfg/env.h"
#include "spdlog/sinks/stdout_color_sinks.h"
#include "spdlog/spdlog.h"

using json = nlohmann::json;

int main(int argc, char *argv[]) {
  // SPDLOG_LEVEL=error,console=info
  spdlog::cfg::load_env_levels();
  // ./nlohmann_json SPDLOG_LEVEL=error,console=info
  spdlog::cfg::load_argv_levels(argc, argv);

  auto console = spdlog::stdout_color_mt("console");

  json j;
  j["project"] = "Bazel Example";
  j["library"] = "nlohmann/json";
  j["version"] = "3.12.0";
  j["status"] = "working";

  spdlog::info("{}", j.dump(2));
  console->info("{}", j.dump(2));

  return 0;
}
