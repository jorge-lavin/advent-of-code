package days

type passport struct {
	birthYear      *int
	issueYear      *int
	expirationYear *int
	heigth         *string
	hairColor      *string
	eyeColor       *string
	passportID     *int
	countryID      *int
}

func (m passport) isValid() bool {
	return m.birthYear != nil &&
		m.issueYear != nil &&
		m.expirationYear != nil &&
		m.heigth != nil &&
		m.hairColor != nil &&
		m.eyeColor != nil &&
		m.passportID != nil

	// countryID is irrelevant if all are fields are present
}

func newPassport(birthYear int, issueYear int, expirationYear int, heigth string, hairColor string, eyeColor string, passportID int, countryID int) *passport {
	return &passport{birthYear: &birthYear, issueYear: &issueYear, expirationYear: &expirationYear, heigth: &heigth, hairColor: &hairColor, eyeColor: &eyeColor, passportID: &passportID, countryID: &countryID}
}

func passportsFromLines(lines []string) []passport {
	return []passport{}
}
