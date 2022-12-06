VERSION := v4.2.1

tag:
	@git tag -a ${VERSION} -m ${VERSION}
	@git push origin ${VERSION}
