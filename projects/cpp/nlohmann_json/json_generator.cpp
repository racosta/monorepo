#include "projects/cpp/nlohmann_json/json_generator.h"
#include <nlohmann/json.hpp>

namespace projects::cpp::nlohmann_json {

std::string GenerateProjectInfo() {
  ::nlohmann::json j;
  j["project"] = "Bazel Example";
  j["library"] = "nlohmann_json";
  j["version"] = "3.11.2";
  j["status"] = "working";

  return j.dump(2);
}

} // namespace projects::cpp::nlohmann_json
