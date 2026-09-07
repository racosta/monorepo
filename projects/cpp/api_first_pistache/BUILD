load("@openapi_tools_generator_bazel//:defs.bzl", "openapi_generator")

openapi_generator(
    name = "user_models_gen",

    # # Restrict execution to model files and core infrastructure only
    # global_property = {
    #     "models": "",             # Generates all schemas under components/schemas
    #     "supportingFiles": "Helpers.h,Helpers.cpp,ModelBase.h,ModelBase.cpp", # Basic serialization helpers
    #     "apis": "false",          # Skip API handlers, controllers, and routing
    #     "apiTests": "false",       # Skip API test stubs
    #     "apiDocs": "false",        # Skip API markdown docs
    #     "modelTests": "false",     # Skip model unit tests
    #     "modelDocs": "false",      # Skip model markdown docs
    # },

    # Generator specific tweaks
    additional_properties = {
        "helpersPackage": "ApiHelpers",
        "hideGenerationTimestamp": "true",  # Recommended for Bazel cache hits
    },
    generator = "cpp-pistache-server",  # You can also use "cpp-qt-client" or custom generators
    spec = "openapi.yaml",
)
