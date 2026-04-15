#include "projects/cpp/rapidjson/json_generator.h"
#include <spdlog/cfg/argv.h>
#include <spdlog/cfg/env.h>
#include <spdlog/sinks/stdout_color_sinks.h>
#include <spdlog/spdlog.h>

int main(int argc, char *argv[]) {
  // SPDLOG_LEVEL=error
  spdlog::cfg::load_env_levels();
  // ./rapidjson SPDLOG_LEVEL=error
  spdlog::cfg::load_argv_levels(argc, argv);

  std::string json_output = projects::cpp::rapidjson::GenerateProjectInfo();
  spdlog::info("{}", json_output);

  return 0;
}
