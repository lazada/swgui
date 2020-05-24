SWAGGER_UI_VERSION := v3.25.4

update:
	curl https://github.com/swagger-api/swagger-ui/blob/$(SWAGGER_UI_VERSION)/dist/swagger-ui-bundle.js -o ./v3/static/swagger-ui-bundle.js
	curl https://github.com/swagger-api/swagger-ui/blob/$(SWAGGER_UI_VERSION)/dist/swagger-ui-standalone-preset.js -o ./v3/static/swagger-ui-standalone-preset.js
	curl https://github.com/swagger-api/swagger-ui/blob/$(SWAGGER_UI_VERSION)/dist/swagger-ui.js -o ./v3/static/swagger-ui.js
	curl https://github.com/swagger-api/swagger-ui/blob/$(SWAGGER_UI_VERSION)/dist/swagger-ui.css -o ./v3/static/swagger-ui.css
	curl https://github.com/swagger-api/swagger-ui/blob/$(SWAGGER_UI_VERSION)/dist/oauth2-redirect.html -o ./v3/static/oauth2-redirect.html
	curl https://github.com/swagger-api/swagger-ui/blob/$(SWAGGER_UI_VERSION)/dist/favicon-32x32.png -o ./v3/static/favicon-32x32.png
	curl https://github.com/swagger-api/swagger-ui/blob/$(SWAGGER_UI_VERSION)/dist/favicon-16x16.png -o ./v3/static/favicon-16x16.png
	go run ./v3/gen/gen.go
