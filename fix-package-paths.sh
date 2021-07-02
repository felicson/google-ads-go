#!/bin/bash

PACKAGES=("common" "enums" "errors" "resources" "services")
ROOT=$1
ADS_VERSION=$2

function fix_package_path() {
    FILE=$1
    PACKAGE=$2
	MATCH="google.golang.org\/genproto\/googleapis\/ads\/googleads\/${ADS_VERSION}\/"
    REPLACE="github.com\/ercling\/google-ads-go\/"

    sed -i "s/$MATCH$PACKAGE/$REPLACE$PACKAGE/g" $FILE
}

echo "fixing packages"
for file in ${ROOT}/google/ads/googleads/${ADS_VERSION}/**/*.pb.go; do
    [ -e "$file" ] || continue
    for p in "${PACKAGES[@]}"; do
        fix_package_path $file $p
    done
done
echo "finished fixing packages"
