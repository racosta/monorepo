"""Rule to create an OCI tar layer from a Java runtime."""

load("@rules_java//java/common:java_common.bzl", "java_common")

def _java_runtime_layer_impl(ctx):
    java_runtime = ctx.attr.java_runtime[java_common.JavaRuntimeInfo]
    output_tar = ctx.actions.declare_file(ctx.label.name + ".tar.gz")

    # Get all Java runtime files
    java_files = java_runtime.files.to_list()

    if not java_files:
        fail("No Java runtime files found in toolchain")

    # Determine version: use explicit version if provided, otherwise auto-detect from label
    version = ctx.attr.version
    if not version or version == "auto":
        # Try to extract version from java_runtime label (e.g., "remotejdk_21" -> "21")
        label_name = ctx.attr.java_runtime.label.name

        # Look for digits in the label name
        version_parts = [c for c in label_name.elems() if c.isdigit()]
        if version_parts:
            version = "".join(version_parts)
        else:
            # Fallback to "21" if we can't detect
            version = "21"

    # Create a manifest of files to include in the tar
    manifest = ctx.actions.declare_file(ctx.label.name + "_manifest.txt")
    manifest_content = []

    # Use JAVA_HOME path from the runtime
    java_home = java_runtime.java_home

    for f in java_files:
        file_path = f.path
        short_path = f.short_path

        # Extract relative path within the JDK installation
        # Handle both "external/remotejdk21_linux/bin/java" and "../remotejdk21_linux/bin/java"
        relative_path = None

        # Try matching against java_home (e.g., "extgernal/remotejdk21_linux")
        if java_home in short_path:
            relative_path = short_path.split(java_home, 1)[1].lstrip("/")
            # Try matching against "../" prefix for external repos

        elif short_path.startswith("../"):
            # Extract repo name from java_home (e.g., "remotejdk21_linux" from "external/remotejdk21_linux")
            repo_name = java_home.split("/")[-1] if "/" in java_home else java_home
            if repo_name in short_path:
                relative_path = short_path.split(repo_name, 1)[1].lstrip("/")

        # Fallback: use basename (should rarely happen with correct logic)
        if not relative_path:
            relative_path = f.basename

        target_path = "usr/lib/jvm/java-{}/{}".format(version, relative_path)
        manifest_content.append("{}={}".format(target_path, file_path))

    ctx.actions.write(
        output = manifest,
        content = "\n".join(manifest_content),
    )

    # Create tar using the manifest
    ctx.actions.run_shell(
        inputs = depset(java_files + [manifest]),
        outputs = [output_tar],
        command = """
      set -euo pipefail
      TAR=tar
      tmpdir=$(mktemp -d)

      # Copy files to tmpdir with proper structure
      while IFS='=' read -r target source; do
        target_dir=$(dirname "$tmpdir/$target")
        mkdir -p "$target_dir"
        if [ -f "$source" ]; then
          cp -L "$source" "$tmpdir/$target"
          # Preserve executable bit for binaries
          if [[ "$target" == */bin/* ]] || [[ "$target" == *.so ]]; then
            chmod +x "$tmpdir/$target" 2>/dev/null || true
          fi
        fi
      done < {manifest}

      # Create tar with specified ownership
      $TAR --owner={user} --group=0 --numeric-owner -czf {output} -C "$tmpdir" .
      rm -rf "$tmpdir"
    """.format(
            manifest = manifest.path,
            output = output_tar.path,
            user = ctx.attr.user,
        ),
        mnemonic = "JavaRuntimeLayer",
    )

    return [DefaultInfo(files = depset([output_tar]))]

java_runtime_layer = rule(
    implementation = _java_runtime_layer_impl,
    attrs = {
        "java_runtime": attr.label(
            default = Label("@rules_java//toolchains:remotejdk_21"),
            providers = [java_common.JavaRuntimeInfo],
        ),
        "version": attr.string(
            default = "auto",
            doc = "Java version number (e.g., '17', '21') for the installation path. Use 'auto' to detect from java_runtime label.",
        ),
        "user": attr.string(
            default = "185",
            doc = "User ID to own the files in the tar (default: '185' for UBI containers).",
        ),
    },
    doc = """Creates a tar.gz layer from a Java runtime for use in OCI images.

    The Java runtime will be installed to /usr/lib/jvm/java-{version}/ in the container.
    Set JAVA_HOME=/usr/lib/jvm/java-{version} in your application to use it.

    The version is automatically detected from the java_runtime label (e.g., "remotejdk_21" -> "21").

    Examples:
      # Java 21 (auto-detected from default)
      java_runtime_layer(
        name = "java_layer",
      )

      # Java 17 (auto-detected from label)
      java_runtime_layer(
        name = "java17_layer",
        java_runtime = "@rules_java//toolchains:remotejdk_17",
      )

      # Explicit version override
      java_runtime_layer(
        name = "java_custom",
        java_runtime = "@rules_java//toolchains:remotejdk_17",
        version = "17",  # Optional: explicity set version
      )

      oci_image(
        name = "image",
        base = "@ubi10-minimal",
        tars = [
          ":java_layer",
          ":app_layer_default.tar.gz",
          # ...
        ],
        env = {
          "JAVA_HOME": "/usr/lib/jvm/java-21",
          "PATH": "/usr/lib/jvm/java-21/bin:$PATH",
        },
      )
    """,
)
