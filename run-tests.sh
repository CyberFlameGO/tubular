#!/bin/bash
# Test the current package under a different kernel.
# Requires virtme and qemu to be installed.

set -eu
set -o pipefail

readonly default_version=5.10.8

if [[ "${1:-}" = "--in-vm" ]]; then
  shift

  mount -t bpf bpf /sys/fs/bpf
  export CCACHE_DISABLE=1
  export GOPATH=/run/go-path
  export GOPROXY=file:///run/go-root/pkg/mod/cache/download
  export GOCACHE=/run/go-cache

  make test
  touch "$1/success"
  exit 0
fi

fetch() {
    echo Fetching "${1}"
    wget -nv -N -P "${2}" "https://unimog.s3.cfdata.org/kernels/${1}"
}

# Pull all dependencies, so that we can run tests without the
# vm having network access.
go mod download

# Use sudo if /dev/kvm isn't accessible by the current user.
sudo=""
if [[ ! -r /dev/kvm || ! -w /dev/kvm ]]; then
  sudo="sudo"
fi
readonly sudo

readonly input="$(mktemp -d)"
readonly output="$(mktemp -d)"
readonly tmp_dir="${TMPDIR:-/tmp}"

if [[ -e "${1:-}" ]]; then
  readonly kernel="${1}"
else
  readonly version="${1:-$default_version}"
  fetch "linux-${version}.bz" "${tmp_dir}"
  readonly kernel="${tmp_dir}/linux-${version}.bz"
fi

echo Testing on "${kernel}"
$sudo virtme-run --kimg "${kernel}" --memory 512M --pwd \
  --rwdir=/run/input="${input}" \
  --rwdir=/run/output="${output}" \
  --rodir=/run/go-path="$(go env GOPATH)" \
  --rwdir=/run/go-cache="$(go env GOCACHE)" \
  --script-sh "PATH=\"$PATH\" $(realpath "$0") --in-vm /run/output" \
  --qemu-opts -smp 2 # need at least two CPUs for some tests

if [[ ! -e "${output}/success" ]]; then
  exit 1
fi

$sudo rm -r "${input}"
$sudo rm -r "${output}"
