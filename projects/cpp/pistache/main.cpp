#include "projects/cpp/fmt_adapters/pistache.h"
#include <pistache/endpoint.h>
#include <pistache/http.h>
#include <pistache/http_defs.h>
#include <pistache/router.h>
#include <spdlog/cfg/argv.h>
#include <spdlog/cfg/env.h>
#include <spdlog/sinks/stdout_color_sinks.h>
#include <spdlog/spdlog.h>

using namespace Pistache;

class HelloHandler : public Http::Handler {
public:
  HTTP_PROTOTYPE(HelloHandler)
  void onRequest(const Http::Request &request,
                 Http::ResponseWriter response) override {
    spdlog::info("Received {} request for {}", request.method(),
                 request.resource());

    response.send(Http::Code::Ok, "Hello, World!");
  }
};

int main(int argc, char *argv[]) {
  // SPDLOG_LEVEL=error
  spdlog::cfg::load_env_levels();
  // ./rapidjson SPDLOG_LEVEL=error
  spdlog::cfg::load_argv_levels(argc, argv);

  constexpr auto port_number = 9080;
  spdlog::info("Starting server on port {}...", port_number);

  Port port(port_number);
  Address addr(Ipv4::any(), port);

  auto opts = Http::Endpoint::options().threads(1);
  Http::Endpoint server(addr);
  server.init(opts);
  server.setHandler(Http::make_handler<HelloHandler>());

  try {
    server.serve();
  } catch (const std::exception &e) {
    spdlog::error("Server error: {}", e.what());
    return 1;
  }

  return 0;
}
