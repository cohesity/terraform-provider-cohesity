#!/bin/bash

CIDR_PREFIX=$1
OUTPUT_FILE=$2

# Function to convert CIDR to subnet mask
cidr_to_subnet_mask() {
    local cidr=$1
    local mask=$((0xFFFFFFFF ^ ((1 << (32 - cidr)) - 1)))
    printf "%d.%d.%d.%d\n" \
      $(( (mask >> 24) & 255 )) \
      $(( (mask >> 16) & 255 )) \
      $(( (mask >> 8) & 255 )) \
      $(( mask & 255 ))
}

# Call the function and write the result to a file
cidr_to_subnet_mask $CIDR_PREFIX > $OUTPUT_FILE
