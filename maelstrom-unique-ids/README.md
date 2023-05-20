brew install openjdk graphviz gnuplot

to run the test cases:
go get github.com/jepsen-io/maelstrom/demo/go
go install .

{two}$ ../maelstrom/maelstrom test -w unique-ids --bin ~/go/bin/maelstrom-unique-ids --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition