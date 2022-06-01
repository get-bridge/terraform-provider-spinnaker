VERSION := v4.2.0

tag:
	@git tag -a ${VERSION} -m ${VERSION}
	@git push origin ${VERSION}
