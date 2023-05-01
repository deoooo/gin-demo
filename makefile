build-doc:
	swag init --parseDependency --parseInternal ./routers/*.go --output ./docs --outputTypes go