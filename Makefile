#!make

include .env

ifdef case
test:
	DOPPLER_KEY=${DOPPLER_KEY} go test -v -run ${case}
else
test:
	DOPPLER_KEY=${DOPPLER_KEY} go test -v
endif
