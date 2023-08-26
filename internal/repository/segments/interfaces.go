package segments

type Repository interface {
	CreateSegment(slug string) error
	DeleteSegment(slug string) (int64, error)
	GetActiveSegmentsIdsBySlugs(slugs []string) (map[string]int64, error)
}
