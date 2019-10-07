package serverv1

const (
	singlePackageTemplate string = `<!doctype html>
<html>
	<head>
		<meta name="go-import" content="{PKGNAME} {VCS} {REPOPATH}">

	</head>
	<body>
		go get {PKGNAME}
	</body>
</html>
`
)

// This might be added. ToDo: figure out why this is needed.
// 		<meta name="go-source" content="{PKGNAME} _ https://sources.dev.pztrn.name/pztrn/giredore/src/branch/master{/dir} https://sources.dev.pztrn.name/pztrn/giredore/src/branch/master{/dir}/{file}#L{line}">
