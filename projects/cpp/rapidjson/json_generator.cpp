#include "projects/cpp/rapidjson/json_generator.h"
#include <rapidjson/document.h>
#include <rapidjson/prettywriter.h>
#include <rapidjson/stringbuffer.h>

namespace projects::cpp::rapidjson {

std::string GenerateProjectInfo() {
  ::rapidjson::Document d;
  d.SetObject();
  auto &allocator = d.GetAllocator();

  d.AddMember("project", "Bazel Example", allocator);
  d.AddMember("library", "rapidjson", allocator);
  d.AddMember("version", "1.1.0", allocator);
  d.AddMember("status", "working", allocator);

  ::rapidjson::StringBuffer buffer;
  ::rapidjson::PrettyWriter<::rapidjson::StringBuffer> writer(buffer);
  writer.SetIndent(' ', 2);
  d.Accept(writer);

  return buffer.GetString();
}

} // namespace projects::cpp::rapidjson
