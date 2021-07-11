# --- Docker ---
docker:
	make docker.pull
	make docker.build

docker.pull:
	docker pull golang:1.16.5-alpine3.14;

# --- Drone ---

drone:
	@stat "cmd/${TARGET}/.drone.yml" > /dev/null && echo "found ${TARGET} .drone.yml"
	drone exec "cmd/${TARGET}/.drone.yml";

# --- Service ---

service:
	make service.docker

service.up.build:
	make service.docker.build

service.up:
	make service.up.build
	docker run -p 8080:8080 lab-less-coffee-${TARGET}

service.docker:
	make service.docker.build

service.docker.build:
	@stat "./cmd/${TARGET}/Dockerfile" > /dev/null && echo "found ${TARGET} Dockerfile"
	docker build -t lab-less-coffee-${TARGET} -f ./cmd/${TARGET}/Dockerfile .


