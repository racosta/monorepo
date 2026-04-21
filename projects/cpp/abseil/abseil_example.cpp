#include "projects/cpp/abseil/abseil_example.h"
#include "absl/strings/str_join.h"

namespace projects::cpp::abseil {

std::string JoinItems(const std::vector<std::string> &items) {
  return absl::StrJoin(items, ", ");
}

} // namespace projects::cpp::abseil
