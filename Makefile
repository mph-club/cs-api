export CURRENT_HEAD = $$(git rev-parse HEAD)

run/server:
	@go run cs-service.go

new-binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./cs-service

#Use this after code push only
docker-build-api:
	#only have to login (the below command) once per 12 hours
	#@eval `aws ecr get-login --region us-east-1 --no-include-email`
	@cp ${GOPATH}/src/mphclub-rest-server/models/models.go ./models/models.go
	@docker build -t cs_portal -f ./Dockerfile .
	@docker tag cs_portal:latest 077003688714.dkr.ecr.us-east-1.amazonaws.com/cs_portal:latest
	@docker tag cs_portal:latest 077003688714.dkr.ecr.us-east-1.amazonaws.com/cs_portal:${CURRENT_HEAD}
	@docker push 077003688714.dkr.ecr.us-east-1.amazonaws.com/cs_portal:latest
	@docker push 077003688714.dkr.ecr.us-east-1.amazonaws.com/cs_portal:${CURRENT_HEAD}
	@kubectl set image deployments/cs-portal-deployment cs-portal=077003688714.dkr.ecr.us-east-1.amazonaws.com/cs_portal:${CURRENT_HEAD}
