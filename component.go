package lily


// Component will deliver a result based on the context
type Component interface {
	Resolve(context interface{}) interface{}
}