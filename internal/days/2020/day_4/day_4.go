package days

import (
	"strings"

	"github.com/lavinj87/advent/internal/lib"
)

type passport struct {
	fields map[string]interface{}
}

func (m passport) isValid() bool {
	return m.fields["byr"] != nil && // birthYear
		m.fields["iyr"] != nil && // issueYear
		m.fields["eyr"] != nil && // expirationYear
		m.fields["hgt"] != nil && // heigth
		m.fields["hcl"] != nil && // hairColor
		m.fields["ecl"] != nil && // eyeColor
		m.fields["pid"] != nil // passportId

	// cid countryID is irrelevant if all are fields are present
}

func (m passport) isValidStrict() bool {
	return m.isValid() &&
		m.isValidBirthYear() &&
		m.isValidIssueYear() &&
		m.isValidExpirationYear() &&
		m.isValidHeigth() &&
		m.isValidHairColor() &&
		m.isValidEyeColor() &&
		m.isValidPassportID()
}

func (m passport) isValidBirthYear() bool {
	return m.isValidYear("byr", 1920, 2002)
}

func (m passport) isValidIssueYear() bool {
	return m.isValidYear("iyr", 2010, 2020)
}

func (m passport) isValidExpirationYear() bool {
	return m.isValidYear("eyr", 2020, 2030)
}
func (m passport) isValidYear(key string, min int, max int) bool {
	yearString := m.fields[key].(string)
	if len(yearString) != 4 {
		return false
	}
	year := lib.StringToInt(yearString)
	return year >= min && year <= max
}

func (m passport) isValidHeigth() bool {
	heigthString := m.fields["hgt"].(string)
	if strings.Contains(heigthString, "cm") {
		parts := strings.Split(heigthString, "cm")
		number := lib.StringToInt(parts[0])
		return number >= 150 && number <= 193
	}
	if strings.Contains(heigthString, "in") {
		parts := strings.Split(heigthString, "in")
		number := lib.StringToInt(parts[0])
		return number >= 59 && number <= 76
	}

	return false
}

func (m passport) isValidHairColor() bool {
	hairColor := m.fields["hcl"].(string)
	if hairColor[0] != '#' {
		return false
	}
	if len(hairColor) != 7 {
		return false
	}
	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	chars := []string{"a", "b", "c", "d", "e", "f"}

	for i := 1; i <= 6; i++ {
		if !(lib.Contains(nums, string(hairColor[i])) || lib.Contains(chars, string(hairColor[i]))) {
			return false
		}
	}
	return true
}

func (m passport) isValidEyeColor() bool {
	return lib.Contains([]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}, m.fields["ecl"].(string))
}

func (m passport) isValidPassportID() bool {
	passportID := m.fields["pid"].(string)
	if len(passportID) != 9 {
		return false
	}

	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, char := range passportID {
		if !lib.Contains(nums, string(char)) {
			return false
		}
	}
	return true
}

func passportsFromLines(lines []string) []passport {
	passports := []passport{}
	for _, passportLines := range lib.Partition(lines, func(line string) bool { return line == "" }) {
		passports = append(passports, passportFromLines(lib.FilterEmpty(passportLines)))
	}
	return passports
}

func passportFromLines(lines []string) passport {
	passport := passport{fields: make(map[string]interface{})}
	for _, line := range lines {
		for _, part := range strings.Split(line, " ") {
			keyValue := strings.Split(part, ":")
			key := keyValue[0]
			value := keyValue[1]
			passport.fields[key] = value
		}
	}
	return passport
}
