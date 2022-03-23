package algebra

// residue field
type ResidueField struct { // implements FiniteField
	r CommutativeRing
	ideal Polynormal // NOTE: shoud we support complet polynormal as a ideal?
}

type Polynormal struct {
	coefficients []int64;
}
