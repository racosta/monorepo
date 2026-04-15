#pragma once

#include <fmt/format.h>
#include <pistache/http_defs.h>

template <> struct fmt::formatter<Pistache::Http::Method> {
  // Parses format specs (e.g., {:x}). We'll just support default.
  constexpr auto parse(format_parse_context &ctx) -> decltype(ctx.begin()) {
    return ctx.begin();
  }

  // Formats the Method enum to the buffer
  template <typename FormatContext>
  auto format(const Pistache::Http::Method &method, FormatContext &ctx) const
      -> decltype(ctx.out()) {
    std::string_view name = "UNKNOWN";
    // There are a bunch of HTTP methods defined in Pistache::Http::Method,
    // but we'll just handle the common ones here.
    // See
    // https://github.com/pistacheio/pistache/blob/master/include/pistache/http_defs.h#L24
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wswitch"
    switch (method) {
    case Pistache::Http::Method::Get:
      name = "GET";
      break;
    case Pistache::Http::Method::Post:
      name = "POST";
      break;
    case Pistache::Http::Method::Put:
      name = "PUT";
      break;
    case Pistache::Http::Method::Delete:
      name = "DELETE";
      break;
    case Pistache::Http::Method::Patch:
      name = "PATCH";
      break;
    case Pistache::Http::Method::Options:
      name = "OPTIONS";
      break;
    }
#pragma clang diagnostic pop
    return fmt::format_to(ctx.out(), "{}", name);
  }
};
