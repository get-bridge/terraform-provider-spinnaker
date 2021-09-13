VERSION := v4.1.0

tag:
	@git tag -a ${VERSION} -m ${VERSION}
	@git push origin ${VERSION}
