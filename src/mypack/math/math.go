package math

/*
* math is the name of a standard library package, but because packages in go
* are hierarchical, we can name ours math within mypack without conflict
*
* import using the full package name, but inside math.go use the basename
*
* call functions in the packge using the basename
* when using two packages with the same basename, an alias can be made
* import math2 "mypack/math"
*
* capitalized functions are publicly accessible from other packages, lower-case
* ones are private (verbosity is reduced by not needing access specifiers)
* generally implementation details are hidden and the api is public
*
* package names match their parent directory name
*/
func Average(xs []float64) float64{
	total := float64(0)
	for _, x := range xs{
		total += x
	}
	return total / float64(len(xs))
}
