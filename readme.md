automatic or algorithmic differentiation in go
----------------------------------------------

a simple mathematical formula language, which is auto-differentiated
and compiled to http://golang.org for high performance.

see https://autodiff.info for live demo.

to get started: make sure you have latest golang.org installed
(e.g. https://golang.org/dl/), or build it yourself via
https://github.com/xoba/goinit

then:

    git clone --recursive https://github.com/xoba/ad.git
    cd ad
    source goinit.sh
    ./install.sh
    run compile -formula "f := sqrt(abs(a+b*b))"
    go run compute.go

for help, you can try:

    run
    run compile -help
    run nn -help

it runs with both scalar and slice variables; e.g.:

    run compile -formula "f:= 2*x[0]+1 + a + x[1] * sin(x[2]) + z/y[0]"

to develop with emacs:

    ./ide.sh

to auto-generate various code:

    lib/gogenerate.sh

to run a simple neural network example:

    run nn

which produces one of these two videos (first one has 5 hidden
units second one has none, and is equivalent to logistic regression):

1. https://s3.amazonaws.com/xoba-videos/test_5.mp4
2. https://s3.amazonaws.com/xoba-videos/test_0.mp4


