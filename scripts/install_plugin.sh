#!/bin/sh -e

# shellcheck disable=SC2002
version="$(cat plugin.yaml | grep "version" | cut -d ' ' -f 2)"
os=$(uname)
echo "Downloading and installing scan v${version} for ${os}..."

url=""
if [ "${os}" = "Linux" ] ; then
    url="../_dist/helm-scan-linux-amd64.tar.gz"
elif [ "${os}" = "Darwin" ] ; then
   url="../_dist/helm-scan-darwin-amd64.tar.gz"
else
    url="../_dist/helm-scan-windows-amd64.tar.gz"
fi

mkdir -p "bin"
mkdir -p "releases/v${version}"

# Download with curl if possible.
# shellcheck disable=SC2230
if [ -x "$(which curl 2>/dev/null)" ]; then
    curl -sSL "${url}" -o "releases/v${version}.tar.gz"
else
    wget -q "${url}" -O "releases/v${version}.tar.gz"
fi
tar xzf "releases/v${version}.tar.gz" -C "releases/v${version}"
if [ "${os}" = "Linux" ] || [ "${os}" = "Darwin" ] ; then
    mv "releases/v${version}/bin/helm-scan" "bin/helm-scan"
else
    mv "releases/v${version}/bin/helm-scan.exe" "bin/helm-scan.exe"
fi
mv "releases/v${version}/plugin.yaml" .