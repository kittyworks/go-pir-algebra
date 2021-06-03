package algebra

type Field interface {
	Add(Field) Field
	Sub(Field) Field
	Mul(Field) Field
	Div(Field) Field
	Neg() Field
	Square() Field
	Sqrt() Field
	Inv() Field
	Exp(Field) Field
	Eq(Field) bool
	Lt(Field) bool
	Gt(Field) bool
	Le(Field) bool
	Ge(Field) bool
	ToString(Field) string
}
