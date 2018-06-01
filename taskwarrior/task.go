package taskwarrior

type task struct {
	ID          int
	Description string
	Urgency     float64
	Status      string
}
