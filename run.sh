#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

usage () {
	echo "USAGE: $0 <day> <part>"
	exit 1
}

main () {
	if [ $# != 2 ]; then
		usage
	fi
	local -ri day="$1"
	local -ri part="$2"

	if [[ -z ${day} || -z ${part} ]]; then
		usage
	fi

	pushd "./days/$(printf '%02d' ${day})" 1>/dev/null
	go run "part${part}.go" "${part}"
	popd 1>/dev/null
}

main "$@"
