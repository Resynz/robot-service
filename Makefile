EXE = robot-service
IMAGE_NAME = $(EXE):0.0.1
build:
	go build -ldflags "-linkmode external -extldflags -static"
docker:
	mv $(EXE) docker-configs
	@echo "building docker image $(IMAGE_NAME) ..."
	docker build -t $(IMAGE_NAME) docker-configs
	rm docker-configs/$(EXE)
run:
	@echo "runing $(EXE) ..."
	docker-compose up -d
stop:
	@echo "stopping $(EXE) ..."
	docker-compose down
clean:
	@echo "cleaning docker image $(IMAGE_NAME) ..."
	docker rmi $(IMAGE_NAME)
log:
	docker-compose logs -f
help:
	@echo "make			-- build $(EXE) application"
	@echo "make docker		-- build docker image"
	@echo "make run 		-- run $(EXE) docker server"
	@echo "make stop		-- stop $(EXE) docker server"
	@echo "make clean 		-- clean docker image"
	@echo "make log 		-- show logs"