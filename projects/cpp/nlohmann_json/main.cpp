#include "projects/cpp/nlohmann_json/json_generator.h"
#include <spdlog/cfg/argv.h>
#include <spdlog/cfg/env.h>
#include <spdlog/sinks/stdout_color_sinks.h>
#include <spdlog/spdlog.h>

int main(int argc, char *argv[]) {
  // SPDLOG_LEVEL=error
  spdlog::cfg::load_env_levels();
  // ./nlohmann_json SPDLOG_LEVEL=error
  spdlog::cfg::load_argv_levels(argc, argv);

  std::string json_output = projects::cpp::nlohmann_json::GenerateProjectInfo();
  spdlog::info("{}", json_output);

  return 0;
}
