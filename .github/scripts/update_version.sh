#!/bin/bash

# # Specify the file path
# file_path="backend/server/metadata.go"

# # Use grep to extract the version string
# version_string=$(grep -o 'Version = "[0-9]\+\.[0-9]\+\.[0-9]\+"' "$file_path")

# # Extract the version number from the version string
# version_number=$(echo "$version_string" | sed 's/Version = "\(.*\)"/\1/')

# echo $version_number

#

# Specify the file paths
source_file="backend/server/metadata.go"
target_file="main.go"

# Use grep to extract the version string from source file
version_string=$(grep -o 'Version = "[0-9]\+\.[0-9]\+\.[0-9]\+"' "$source_file")

# Extract the version number from the version string
version_number=$(echo "$version_string" | sed 's/Version = "\(.*\)"/\1/')

echo "Extracted version: $version_number"

# Update target file with the extracted version number
sed -i "s#// @version[[:space:]]\+[0-9]\+\.[0-9]\+\.[0-9]\+#// @version\t\t$version_number#" "$target_file"

echo "Updated target file with new version: $version_number"
