package algebra

type Field interface {
	// returns result of a + b
	Add(Field) Field

	// returns result of a - b
	Sub(Field) Field

	// returns result of a * b
	Mul(Field) Field

	// returns result of a / b
	Div(Field) Field

	// returns -a
	Neg() Field

	// returns root of a
	Square() Field

	// returns square of a
	Sqrt() Field

	// returns 1/a
	Inv() Field

	// returns a ^ b
	Exp(Field) Field

	// returns true if a = b
	Eq(Field) bool

	// returns true if a < b
	Lt(Field) bool

	// returns true if a > b
	Gt(Field) bool

	// returns true if a <= b
	Le(Field) bool

	// returns true if a >= b
	Ge(Field) bool

	// convert field to a string to display
	ToString(Field) string

	// returns a indetity element of addition
	GetAddElement() Field

	// returns a indetity element of multiply
	GetMulElement() Field
}
