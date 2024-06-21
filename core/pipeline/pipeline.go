package pipeline

type IPipeline interface {
	Name() string
	Run(args Arguments) error
}
