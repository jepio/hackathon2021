.PHONY: all deploy
all:
	CGO_ENABLED=0 go build -v

deploy: all
	func azure functionapp publish clctranspiler
