#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SOLUTIONS_DIR="$ROOT/solutions"

usage() {
  cat <<'EOF'
Usage:
  ./scripts/newday.sh            # create next day based on existing solutions/dayXX
  ./scripts/newday.sh 7          # create day07
  ./scripts/newday.sh 07         # create day07

Notes:
- Creates: solutions/dayXX/{dayXX.go,dayXX_test.go,example.txt,input.txt}
- Uses module path from go.mod for imports
EOF
}

if [[ "${1:-}" == "-h" || "${1:-}" == "--help" ]]; then
  usage
  exit 0
fi

# Determine module path from go.mod
if [[ ! -f "$ROOT/go.mod" ]]; then
  echo "Error: go.mod not found at repo root: $ROOT"
  exit 1
fi
MODULE_PATH="$(awk 'NR==1 && $1=="module" {print $2}' "$ROOT/go.mod")"
if [[ -z "${MODULE_PATH}" ]]; then
  echo "Error: could not determine module path from go.mod"
  exit 1
fi

mkdir -p "$SOLUTIONS_DIR"

# Determine day
day=""
if [[ -n "${1:-}" ]]; then
  if [[ ! "$1" =~ ^[0-9]{1,2}$ ]]; then
    echo "Error: day must be 1..25 (got: $1)"
    exit 1
  fi
  day="$1"
else
  # Infer next day from existing folders like day01, day02, ...
  last="$(ls -1 "$SOLUTIONS_DIR" 2>/dev/null | grep -E '^day[0-9]{2}$' | sort | tail -n1 || true)"
  if [[ -z "$last" ]]; then
    day="1"
  else
    last_num="${last#day}"
    day=$((10#$last_num + 1))
  fi
fi

if (( day < 1 || day > 25 )); then
  echo "Error: day must be in range 1..25 (got: $day)"
  exit 1
fi

DAY_PADDED="$(printf "%02d" "$day")"
DAY_DIR="$SOLUTIONS_DIR/day$DAY_PADDED"

if [[ -e "$DAY_DIR" ]]; then
  echo "Error: $DAY_DIR already exists"
  exit 1
fi

mkdir -p "$DAY_DIR"

GO_FILE="$DAY_DIR/day$DAY_PADDED.go"
TEST_FILE="$DAY_DIR/day${DAY_PADDED}_test.go"
EXAMPLE_FILE="$DAY_DIR/example.txt"
INPUT_FILE="$DAY_DIR/input.txt"

cat > "$GO_FILE" <<EOF
package day$DAY_PADDED

import "$MODULE_PATH/internal/aoc"

type sol struct{}

func (sol) Day() int { return $((10#$DAY_PADDED)) }

func (sol) Part1(input string) (any, error) {
	// TODO: parse + solve
	return nil, nil
}

func (sol) Part2(input string) (any, error) {
	// TODO: parse + solve
	return nil, nil
}

func init() { aoc.Register(sol{}) }
EOF

cat > "$TEST_FILE" <<'EOF'
package dayXX

import (
	"os"
	"path/filepath"
	"testing"
)

func mustReadFile(t *testing.T, name string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(".", name))
	if err != nil {
		t.Fatalf("read %s: %v", name, err)
	}
	return string(b)
}

func TestPart1Example(t *testing.T) {
	input := mustReadFile(t, "example.txt")
	got, err := sol{}.Part1(input)
	if err != nil {
		t.Fatalf("Part1 error: %v", err)
	}
	_ = got
	// TODO: set expected
	// if got != 123 { t.Fatalf("got %v", got) }
}

func TestPart2Example(t *testing.T) {
	input := mustReadFile(t, "example.txt")
	got, err := sol{}.Part2(input)
	if err != nil {
		t.Fatalf("Part2 error: %v", err)
	}
	_ = got
	// TODO: set expected
	// if got != 456 { t.Fatalf("got %v", got) }
}
EOF

# Replace placeholder package name dayXX -> dayNN
# (portable across macOS/Linux)
perl -0777 -i -pe "s/package dayXX/package day$DAY_PADDED/g" "$TEST_FILE"

cat > "$EXAMPLE_FILE" <<'EOF'
# Paste the example input here
EOF

: > "$INPUT_FILE"

echo "✅ Created:"
echo "  $DAY_DIR/"
echo "    - day$DAY_PADDED.go"
echo "    - day${DAY_PADDED}_test.go"
echo "    - example.txt"
echo "    - input.txt"
echo

