build-doc:
	swag init --parseDependency --parseInternal ./routers/*.go --output ./docs --outputTypes go

build:
	docker build -t gin_demo .