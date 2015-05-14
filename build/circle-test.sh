#!/bin/bash

set -e

outdir="${TMPDIR}"
if [ "${CIRCLE_ARTIFACTS}" != "" ]; then
    outdir="${CIRCLE_ARTIFACTS}"
fi

build=$(dirname $0)/circle-build.sh

time ${build} make check | tee "${outdir}/check.log"; test ${PIPESTATUS[0]} -eq 0

time ${build} /bin/bash -c "(go generate ./... && git ls-files --modified --deleted --others --exclude-standard | diff /dev/null -) || (git add -A && git diff -u HEAD && false)" | tee "${outdir}/generate.log"; test ${PIPESTATUS[0]} -eq 0

match='^panic|^[Gg]oroutine \d+|(read|write) by.*goroutine|DATA RACE'
time ${build} make testrace \
     RACETIMEOUT=5m TESTFLAGS='-v -vmodule=multiraft=5,raft=1' | \
    tee "${outdir}/testrace.log" | \
    grep -E "^\--- (PASS|FAIL)|${match}"

if [ "${CIRCLE_TEST_REPORTS}" != "" ]; then
    if [ -f "${outdir}/testrace.log" ]; then
        mkdir -p "${CIRCLE_TEST_REPORTS}/race"
	${build} go2xunit < "${outdir}/testrace.log" \
	       > "${CIRCLE_TEST_REPORTS}/race/testrace.xml"
    fi
fi

find "${outdir}" -name '*.log' -type f -exec \
     grep -B 5 -A 10 -E "^\-{0,3} *FAIL|${match}" {} ';' > "${outdir}/excerpt.txt"
test ! -s "${outdir}/excerpt.txt"
