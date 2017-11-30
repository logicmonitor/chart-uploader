NAMESPACE  := logicmonitor
REPOSITORY := chart-uploader
VERSION    := 0.1.0-alpha.0

all:
	docker build --build-arg VERSION=$(VERSION) --build-arg CI=$(CI) -t $(NAMESPACE)/$(REPOSITORY):latest .
	docker tag $(NAMESPACE)/$(REPOSITORY):latest $(NAMESPACE)/$(REPOSITORY):$(VERSION)
