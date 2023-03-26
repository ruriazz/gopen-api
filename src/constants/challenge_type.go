package constants

type ChallengeType int

const (
	DNS ChallengeType = iota
	ACME
)

func (c ChallengeType) String() string {
	return [...]string{"dns", "acme"}[c]
}

func (c ChallengeType) LongString() string {
	return [...]string{"DNS Record", "ACME"}[c]
}
