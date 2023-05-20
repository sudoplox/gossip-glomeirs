brew install openjdk graphviz gnuplot

to run the test cases:
go get github.com/jepsen-io/maelstrom/demo/go
go install .

{maelstrom-echo}$ ../maelstrom/maelstrom test -w echo --bin ~/go/bin/maelstrom-echo --node-count 1 --time-limit 10