VERSION := <VERSION>
ACCOUNT := <ACCOUNT_ID>
REPO := <REPO_NAME>

publish:
	docker build -f server/Dockerfile . -t kiss-server
	docker tag kiss-server ${ACCOUNT}.dkr.ecr.eu-west-1.amazonaws.com/${REPO}/kiss:${VERSION}
	docker push  ${ACCOUNT}.dkr.ecr.eu-west-1.amazonaws.com/${REPO}/kiss:${VERSION}