package pipeline

type IPipeline interface {
	Name() string
	Args() Arguments
	Run(args Arguments) error
}
