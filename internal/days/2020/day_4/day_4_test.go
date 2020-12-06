package days

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/lavinj87/advent/internal/lib"
)

func Test_isValidPassport(t *testing.T) {
	tests := []struct {
		passport passport
		expected bool
	}{{
		passport: passport{},
		expected: false,
	}, {
		// all fields
		passport: passport{fields: map[string]interface{}{"byr": "", "iyr": "", "eyr": "", "hgt": "", "hcl": "", "ecl": "", "pid": "", "cid": ""}},
		expected: true,
	}, {
		// all except countryId
		passport: passport{fields: map[string]interface{}{"byr": "", "iyr": "", "eyr": "", "hgt": "", "hcl": "", "ecl": "", "pid": ""}},
		expected: true,
	}}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			lib.AssertEquals(t, tt.expected, (tt.passport.isValid()))
		})
	}
}

func Test_readPassports(t *testing.T) {
	lines := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	}

	expected := []passport{
		passport{fields: map[string]interface{}{"byr": "1937", "iyr": "2017", "eyr": "2020", "hgt": "183cm", "hcl": "#fffffd", "ecl": "gry", "pid": "860033327", "cid": "147"}},
		passport{fields: map[string]interface{}{"byr": "1929", "iyr": "2013", "eyr": "2023", "hcl": "#cfa07d", "ecl": "amb", "pid": "028048884", "cid": "350"}},
		passport{fields: map[string]interface{}{"byr": "1931", "iyr": "2013", "eyr": "2024", "hgt": "179cm", "hcl": "#ae17e1", "ecl": "brn", "pid": "760753108"}},
		passport{fields: map[string]interface{}{"iyr": "2011", "eyr": "2025", "hgt": "59in", "hcl": "#cfa07d", "ecl": "brn", "pid": "166559648"}},
	}

	actual := passportsFromLines(lines)

	if !cmp.Equal(expected, actual, cmp.AllowUnexported(passport{})) {
		t.Errorf("Expected:%v. Actual:%v", expected, actual)
	}
}

func Test_isValidFromLines(t *testing.T) {
	lines := []string{}
	lib.Lines("day_4_input.txt", func(line string) {
		lines = append(lines, line)
	})

	passports := passportsFromLines(lines)
	validPassports := 0
	for _, passport := range passports {
		if passport.isValid() {
			validPassports++
		}
	}

	expected := 182
	if expected != validPassports {
		t.Errorf("Expected:%v valid passports, got:%v", expected, validPassports)

	}
}

func Test_isValidStrict(t *testing.T) {
	tests := []struct {
		passport passport
		expected bool
	}{{
		passport: passport{fields: map[string]interface{}{"eyr": "1972", "cid": "100", "hcl": "#18171d", "ecl": "amb", "hgt": "170", "pid": "186cm", "iyr": "2018", "byr": "1926"}},
		expected: false,
	}, {
		passport: passport{fields: map[string]interface{}{"pid": "087499704", "hgt": "74in", "ecl": "grn", "iyr": "2012", "eyr": "2030", "byr": "1980", "hcl": "#623a2f"}},
		expected: true,
	}}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			lib.AssertEquals(t, tt.expected, (tt.passport.isValidStrict()))
		})
	}
}

func Test_isValidStrictFromLines(t *testing.T) {
	lines := []string{}
	lib.Lines("day_4_input.txt", func(line string) {
		lines = append(lines, line)
	})

	passports := passportsFromLines(lines)
	validPassports := 0
	for _, passport := range passports {
		if passport.isValidStrict() {
			validPassports++
		}
	}

	expected := 109
	if expected != validPassports {
		t.Errorf("Expected:%v valid passports, got:%v", expected, validPassports)

	}
}
