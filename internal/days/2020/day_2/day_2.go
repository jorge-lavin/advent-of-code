package days

import (
	"strings"

	"github.com/lavinj87/advent/internal/lib"
)

func day2(inputPath string, validator func(record *passwordRecord) bool) int {
	var validPasswordsCounter = 0

	lib.Lines(inputPath, func(line string) {
		record := newPasswordRecord(line)
		if validator(record) {
			validPasswordsCounter = validPasswordsCounter + 1
		}
	})

	return validPasswordsCounter
}

// isValidPassword returns wether the password is valid according to the first policy
func isValidPasswordByCount(record *passwordRecord) bool {
	count := strings.Count(record.password, record.repeated)

	return count >= record.min && count <= record.max
}

// isValidPasswordByPosition returns wether the password is valid according to the first polic
func isValidPasswordByPosition(record *passwordRecord) bool {
	isFirstMached := string(record.password[record.min-1]) == record.repeated
	isSecondMatched := string(record.password[record.max-1]) == record.repeated

	if isFirstMached {
		return !isSecondMatched
	}

	return isSecondMatched

}

type passwordRecord struct {
	min      int
	max      int
	repeated string
	password string
}

func newPasswordRecord(line string) *passwordRecord {
	// 4-8 n: dnjjrtclnzdnghnbnn
	parts := strings.Split(line, " ")
	record := passwordRecord{}
	// 4-8
	minMaxParts := strings.Split(parts[0], "-")
	record.min = lib.StringToInt(minMaxParts[0])
	record.max = lib.StringToInt(minMaxParts[1])
	// n:
	repeatedParts := strings.Split(parts[1], ":")
	record.repeated = repeatedParts[0]
	record.password = parts[2]

	return &record
}
