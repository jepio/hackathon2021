.PHONY: all deploy
all:
	go build

deploy: all
	func azure functionapp publish clctranspiler
