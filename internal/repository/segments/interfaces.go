package segments

//go:generate mockgen -destination=mocks.go -package=$GOPACKAGE -source=interfaces.go

type Repository interface {
	CreateSegment(slug string, percentageUsers int64) error
	DeleteSegment(slug string) (bool, error)
	GetActiveSegmentsIdsBySlugs(slugs []string) (map[string]int64, error)
	GetPercentSegments() (map[string]Segment, error)
}
