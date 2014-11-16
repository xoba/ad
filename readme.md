automatic or algorithmic differentiation
----------------------------------------

strategy: mini-language (i.e., a dsl) for functions 
and formulas, then compile it to go, for high performance.

to get started: make sure you have latest golang.org installed; e.g. https://golang.org/dl/ --- or,
build it yourself via https://github.com/xoba/goinit

then:

    git clone --recursive  git@github.com:xoba/ad.git
    cd ad
    source goinit.sh
    ./install.sh
    run parse -formula "sqrt(a+b*b)"
    go run compute.go

for help, you can try:

    run
    run parse -help

