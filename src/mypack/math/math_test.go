package math

import(
    "testing"
)

type testpair struct{
    values []float64
    average float64
}

var tests = []testpair{
    { []float64{1, 2}, 1.5 },
    { []float64{1, 1, 1, 1}, 1 },
    { []float64{4, 2}, 3 },
}

/*
 * Using the testing package, we can define test cases as functions named Test*
 * accepting a pointer to testing.T. Error conditions are indicated using
 * t.Error. Run the test file with go test.
 */
func TestAverage(t *testing.T){
    v := Average([]float64{1, 2})
    if v != 1.5{
        t.Error("expected 1.5, for ", v)
    }
}

/*
 * Test multiple slices of values. Note that slice.values expands the elements
 * so that they can be passed to a variadic function.
 */
func TestAverageMultiple(t *testing.T){
    for _, pair := range tests{
        // expands tuple
        v := Average(pair.values)

        if v != pair.average{
            t.Error("for", pair.values, "expected", pair.average, "got", v)
        }
    }
}


