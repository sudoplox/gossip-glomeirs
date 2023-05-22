brew install openjdk graphviz gnuplot

to run the test cases:
go get github.com/jepsen-io/maelstrom/demo/go
go install .

{maelstrom-echo}$ ../maelstrom/maelstrom test -w broadcast --bin ~/go/bin/maelstrom-broadcast-3a --node-count 1 --time-limit 20 --rate 10