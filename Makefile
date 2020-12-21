.PHONY: all

all: toc2rgb kelvincalc

toc2rgb:
	go build -o toc2rgb ./cmd/toc2rgb/

kelvincalc:
	go build -o kelvincalc ./cmd/kelvincalc/