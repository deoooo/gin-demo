build-doc:
	swag init --parseDependency --parseInternal ./routers/*.go --output ./docs --outputTypes go

build:
	docker build -t aa864451000/gin_demo:latest .
	docker push aa864451000/gin_demo:latest