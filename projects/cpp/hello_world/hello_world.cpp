#include "greeter.h"
#include <spdlog/cfg/argv.h>
#include <spdlog/cfg/env.h>
#include <spdlog/sinks/stdout_color_sinks.h>
#include <spdlog/spdlog.h>

int main(int argc, char *argv[]) {
  // SPDLOG_LEVEL=error,console=info
  spdlog::cfg::load_env_levels();
  // ./hello_world SPDLOG_LEVEL=error,console=info
  spdlog::cfg::load_argv_levels(argc, argv);

  auto console = spdlog::stdout_color_mt("console");

  auto greeting = greet();

  spdlog::info(greeting);
  console->info(greeting);

  return 0;
}
