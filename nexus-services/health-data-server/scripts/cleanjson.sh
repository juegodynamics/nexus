#!/bin/bash

# Loop through all JSON files in the current directory
for file in health-data-server/data/*.json; do
  # Extract the UUID part of the filename using sed
  uuid=$(echo "$file" | sed -n 's/.*_\([0-9a-fA-F-]\{36\}\)\.json/\1/p')
  # Check if uuid is not empty
  if [ -n "$uuid" ]; then
    # Rename the file to just the UUID
    mv "$file" "health-data-server/data/$uuid.json"
  else
    echo "No valid UUID found in filename: $file"
  fi
done

echo "Renaming complete."
