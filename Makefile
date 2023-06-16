VERSION := v4.2.2

tag:
	@git tag -a ${VERSION} -m ${VERSION}
	@git push origin ${VERSION}
