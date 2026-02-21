"""Macro to create py_iamge_layer outputs with non-root ownership.

This fixes permission issues when running Bazel Python binaries in UBI containers
as non-root users. The py_image_layer rule generates tar archives with root ownership
which prevents Python from creating venvs or writing bytecode cache files.
"""

load("@aspect_rules_py//py:defs.bzl", "py_image_layer")

def py_image_layer_owned(
        name,
        binary,
        user = "185",
        **kwargs):
    """Create py_image_layer outputs with specified user ownership.

    This macro wraps py_image_layer and post-processes the output tars to change
    file ownership from root (0:0) to a specified non-root user. This is necessary
    for running Python service in UBI-based containers as non-root users.

    Args:
      name: Name of this target (will be used to generate the py_image_layer name)
      binary: The py_binary target to package
      user: The numeric user ID to own the files (default: "185" for UBI containers)
      **kwargs: Additional arguments passed to py_image_layer

    Outputs:
      Three tar.gz files with ownership changed to specified user:
      - {name}_default.tar.gz
      - {name}_packages.tar.gz
      - {name}_interpreter.tar.gz

    Example:
      ```
      py_image_layer_owned(
        name = "app_layer",
        binary = ":main_app",
        user = "185",
      )

      oci_image(
        name = "image",
        tars = [
          ":app_layer_default.tar.gz",
          ":app_layer_packages.tar.gz",
          ":app_layer_interpreter.tar.gz",
        ],
        user = "185",
      )
      ```
    """

    base_layer_name = name + "_base"
    py_image_layer(
        name = base_layer_name,
        binary = binary,
        **kwargs
    )

    native.genrule(
        name = name,
        srcs = [":" + base_layer_name],
        outs = [
            name + "_default.tar.gz",
            name + "_packages.tar.gz",
            name + "_interpreter.tar.gz",
        ],
        cmd = """
          # Process each layer and change ownership to {user}:0
          for src in $(SRCS); do
            layer_name=$$(basename $$src)
            case "$$layer_name" in
              *default*)
                out=$(location {name}_default.tar.gz)
                ;;
              *packages*)
                out=$(location {name}_packages.tar.gz)
                ;;
              *interpreter*)
                out=$(location {name}_interpreter.tar.gz)
                ;;
              *)
                continue
                ;;
            esac

            # Extract, change ownership, repackage
            tmpdir=$$(mktemp -d)
            tar -xzf $$src -C $$tmpdir
            tar --owner={user} --group=0 -czf $$out -C $$tmpdir .
            rm -rf $$tmpdir
          done
        """.format(
            name = name,
            user = user,
        ),
        tags = ["no-remote"],
        visibility = ["//visibility:public"],
    )
