
# image settings
image := "hview"
reg := "registry.gross-bau.de"
# get the current date as unix timestamp
tag := `date +%s`

# deployment settings
namespace := "test"
deployment := "hview-deployment"
containername := "hview"
containerimage :=  reg + "/" + image + ":"  + tag
containerimage_latest :=  reg + "/" + image + ":"  + "latest"
# tasks

build:
	docker build --platform linux/amd64 -t {{image}}:{{tag}} .

run: build
	docker run -it -p "8080:8080" "{{image}}:{{tag}}"

push: build
	docker tag {{image}}:{{tag}} {{containerimage}}
	docker tag {{image}}:{{tag}} {{containerimage_latest}}
	docker push {{containerimage}}
	docker push {{containerimage_latest}}

deploy: push
	kubectl -n {{namespace}} set image deployment/{{deployment}} {{containername}}={{containerimage}}
	kubectl -n {{namespace}}  rollout status deployment/{{deployment}}
