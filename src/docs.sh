#!/bin/bash

# prints documentation for one function, the comment immediately above the
# function definition is used
godoc mypack/math Average

# serve all installed package documentation over http
godoc -http=":8000"
