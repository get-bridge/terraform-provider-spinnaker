VERSION := v4.1.1

tag:
	@git tag -a ${VERSION} -m ${VERSION}
	@git push origin ${VERSION}
