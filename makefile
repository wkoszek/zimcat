ZIM_URL=https://dumps.wikimedia.org/kiwix/zim/wikipedia/wikipedia_en_mathematics_mini_2022-07.zim
ZIM_FN=$(shell /bin/ls -1 `pwd`/*.zim)

zimcat: zimcat.go
	go build zimcat.go
	GOOS=linux go build -o zimcat.linux zimcat.go

dow:
	curl -o $(ZIM_FN) $(ZIM_URL)

clean:
	rm -rf zimcat

test:
	./zimcat $(ZIM_FN)
