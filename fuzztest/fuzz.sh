# go get -u github.com/dvyukov/go-fuzz/go-fuzz-build
# go get -u github.com/dvyukov/go-fuzz/go-fuzz
go-fuzz-build github.com/judwhite/talks-gofuzz/fuzztest
go-fuzz -bin=./fuzztest-fuzz.zip -workdir=fuzz
