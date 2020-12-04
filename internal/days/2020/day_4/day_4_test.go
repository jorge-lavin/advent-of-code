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
		passport: passport{birthYear: new(int), issueYear: new(int), expirationYear: new(int), heigth: new(string), hairColor: new(string), eyeColor: new(string), passportID: new(int), countryID: new(int)},
		expected: true,
	}, {
		passport: passport{birthYear: new(int), issueYear: new(int), expirationYear: new(int), heigth: new(string), hairColor: new(string), eyeColor: new(string), passportID: new(int)},
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

	expected := []*passport{
		newPassport(1937, 2017, 2020, "183cm", "#fffffd", "gry", 860033327, 147),
		newPassport(1937, 2017, 2020, "183cm", "#fffffd", "gry", 860033327, 147),
		newPassport(1937, 2017, 2020, "183cm", "#fffffd", "gry", 860033327, 147),
		newPassport(1937, 2017, 2020, "183cm", "#fffffd", "gry", 860033327, 147),
	}

	actual := passportsFromLines(lines)

	if !cmp.Equal(expected, actual) {
		t.Errorf("Expected:%v. Actual:%v", expected, actual)
	}
}
