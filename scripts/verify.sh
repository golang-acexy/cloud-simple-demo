#!/usr/bin/env bash

set -euo pipefail

BASE_URL="${BASE_URL:-http://127.0.0.1:8080}"
USER_A=10001
USER_B=20002
APP_PID=""

cleanup() {
    if [[ -n "${APP_PID}" ]]; then
        kill "${APP_PID}" 2>/dev/null || true
        wait "${APP_PID}" 2>/dev/null || true
    fi
}
trap cleanup EXIT

fail() {
    echo "FAILED: $1" >&2
    exit 1
}

assert_contains() {
    local body="$1"
    local expected="$2"
    local message="$3"
    [[ "${body}" == *"${expected}"* ]] || fail "${message}: ${body}"
}

extract_id() {
    printf '%s' "$1" | sed -n 's/.*"data":\([0-9][0-9]*\).*/\1/p'
}

request() {
    local method="$1"
    local path="$2"
    local body="${3:-}"
    local user_id="${4:-}"
    local args=(-sS -X "${method}" "${BASE_URL}/${path}")
    if [[ -n "${body}" ]]; then
        args+=(-H "Content-Type: application/json" -d "${body}")
    fi
    if [[ -n "${user_id}" ]]; then
        args+=(-H "X-User-ID: ${user_id}")
    fi
    curl "${args[@]}"
}

echo "Starting cloud-simple-demo..."
GOMODCACHE="${GOMODCACHE:-/Users/acexy/Repository/cache/golang}" go run ./cmd >"${TMPDIR:-/tmp}/cloud-simple-demo.log" 2>&1 &
APP_PID=$!

for _ in {1..60}; do
    if curl -sS "${BASE_URL}/adm/department/query" -H "Content-Type: application/json" -d '{}' >/dev/null 2>&1; then
        break
    fi
    sleep 0.5
done
kill -0 "${APP_PID}" 2>/dev/null || fail "application failed to start; see ${TMPDIR:-/tmp}/cloud-simple-demo.log"

suffix="$(date +%s)"
department_code="D-${suffix}"
employee_no="E-${suffix}"

response="$(request POST adm/department/save "{\"code\":\"${department_code}\",\"name\":\"Platform\",\"description\":\"framework verification\",\"status\":1}")"
assert_contains "${response}" '"statusCode":200' "create department"
department_id="$(extract_id "${response}")"
[[ -n "${department_id}" ]] || fail "department ID missing: ${response}"

response="$(request POST usr/employee/save "{\"departmentId\":${department_id},\"userId\":99999,\"employeeNo\":\"${employee_no}\",\"name\":\"Alice\",\"email\":\"alice@example.com\",\"status\":1}" "${USER_A}")"
assert_contains "${response}" '"statusCode":200' "create employee"
employee_id="$(extract_id "${response}")"
[[ -n "${employee_id}" ]] || fail "employee ID missing: ${response}"

response="$(request GET "usr/employee/by-id/${employee_id}" "" "${USER_A}")"
assert_contains "${response}" '"userId":10001' "authority must override submitted userId"
assert_contains "${response}" "\"employeeNo\":\"${employee_no}\"" "owner should query employee"

response="$(request GET "usr/employee/by-id/${employee_id}" "" "${USER_B}")"
[[ "${response}" != *"${employee_no}"* ]] || fail "another user accessed protected employee: ${response}"

response="$(request POST usr/employee/query "{\"name\":\"Alice\"}" "${USER_A}")"
assert_contains "${response}" "${employee_no}" "owner list query"

response="$(request POST usr/employee/query-by-page "{\"number\":1,\"size\":10,\"condition\":{\"name\":\"Alice\"},\"timeRanges\":[{\"field\":\"createdAt\",\"start\":1704067200000}]}" "${USER_A}")"
assert_contains "${response}" '"total":1' "owner page query"

response="$(request PUT "usr/employee/by-id/${employee_id}" "{\"name\":\"Alice Updated\",\"userId\":99999}" "${USER_A}")"
assert_contains "${response}" '"statusCode":200' "owner update"

response="$(request DELETE "usr/employee/by-id/${employee_id}" "" "${USER_B}")"
assert_contains "${response}" '"statusCode":400' "another user delete must fail"

response="$(request DELETE "usr/employee/by-id/${employee_id}" "" "${USER_A}")"
assert_contains "${response}" '"statusCode":200' "owner delete"

response="$(request DELETE "adm/department/by-id/${department_id}")"
assert_contains "${response}" '"statusCode":200' "delete department"

echo "PASS: generator, GORM repository, cloud web CRUD, authority isolation, and time-range pagination"
