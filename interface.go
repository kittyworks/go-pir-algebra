package algebra

type Group interface {
	// returns result of a + b
	Add(Field) Field

	// returns result of a - b
	Sub(Field) Field

	// returns -a
	Neg() Field

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
}

type Ring interface {
	Group

	// returns result of a * b
	Mul(Field) Field

	// returns a ^ b
	Exp(Field) Field

	// returns a * a
	Square() Field
}

type CommutativeRing interface {
	Ring
}

type IntegralDomain interface {
	CommutativeRing

	// returns result of a / b
	Div(Field) Field

	// returns 1/a
	Inv() Field

	// returns root of a
	Sqrt() Field
}

type Field interface {
	IntegralDomain

	// returns a indetity element of multiply
	GetMulElement() Field
}

type FiniteField interface {
	Field
}
