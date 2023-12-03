# https://gist.github.com/prwhite/8168133
help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

check-aoc-cookie:  ## ensures $AOC_SESSION_COOKIE env var is set
	@ test $${AOC_SESSION_COOKIE?env var not set}

input: check-aoc-cookie ## get input, requires $AOC_SESSION_COOKIE, optional: $DAY and $YEAR
	@ if [[ -n $$DAY && -n $$YEAR ]]; then \
		go run input/main.go -day $(DAY) -year $(YEAR) -cookie $(AOC_SESSION_COOKIE); \
	elif [[ -n $$DAY ]]; then \
		go run input/main.go -day $(DAY) -cookie $(AOC_SESSION_COOKIE); \
	else \
		go run input/main.go -cookie $(AOC_SESSION_COOKIE); \
	fi