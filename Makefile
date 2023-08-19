#!make

include .env

ifdef arg
test:
	DOPPLER_KEY=${DOPPLER_KEY} go test -v -run ${arg}
else
test:
	DOPPLER_KEY=${DOPPLER_KEY} go test -v
endif
