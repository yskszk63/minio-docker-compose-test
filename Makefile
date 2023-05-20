.PHONY: prepare
prepare:
	mc alias set s3 http://s3:9000 $$AWS_ACCESS_KEY_ID $$AWS_SECRET_ACCESS_KEY
	mc mb -p s3/test s3/test2
	echo 'Hello, World!' | aws s3 cp --endpoint-url http://s3:9000 - s3://test/hello.txt

.PHONY: run
run:
	go run .
