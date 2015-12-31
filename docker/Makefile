NAME = artifact.stage.shutterfly.com/monitoring/aion

ifndef CHANGESET
       CHANGESET=local
endif

TAG = $(CHANGESET)

.PHONY: build deploy

build:
	@echo "*** building $(NAME):$(TAG) ***"
	docker build --no-cache=true --rm --tag $(NAME):$(TAG) ./prod/

deploy: build
	@echo "*** deploying $(NAME):$(TAG) ***"
	docker push $(NAME):$(TAG)
