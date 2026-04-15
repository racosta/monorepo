#include <rapidjson/document.h>
#include <rapidjson/prettywriter.h>
#include <rapidjson/stringbuffer.h>
#include <spdlog/cfg/argv.h>
#include <spdlog/cfg/env.h>
#include <spdlog/sinks/stdout_color_sinks.h>
#include <spdlog/spdlog.h>

int main(int argc, char *argv[]) {
  // SPDLOG_LEVEL=error
  spdlog::cfg::load_env_levels();
  // ./rapidjson SPDLOG_LEVEL=error
  spdlog::cfg::load_argv_levels(argc, argv);

  rapidjson::Document d;
  d.SetObject();
  rapidjson::Document::AllocatorType &allocator = d.GetAllocator();

  d.AddMember("project", "Bazel Example", allocator);
  d.AddMember("library", "rapidjson", allocator);
  d.AddMember("version", "1.1.0", allocator);
  d.AddMember("status", "working", allocator);

  rapidjson::StringBuffer buffer;
  rapidjson::PrettyWriter<rapidjson::StringBuffer> writer(buffer);
  writer.SetIndent(' ', 2);
  d.Accept(writer);

  spdlog::info("{}", buffer.GetString());

  return 0;
}
