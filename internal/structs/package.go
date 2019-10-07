package structs

// Package defines structure for 'pkg set' request and for storing it's
// data in configuration.
type Package struct {
	// Description is an additional and optional description that
	// can be show on package's page.
	Description string
	// OriginalPath is a package original path without domain part.
	// E.g. for package "go.example.tld/group/pkgname" you should
	// put here "/group/pkgname".
	OriginalPath string
	// RealPath is a path where package will be found. It should
	// contain VCS path, e.g. "https://github.com/user/project.git".
	RealPath string
	// VCS is a versioning control system used for package. Everything
	// that is supported by "go get" is applicable.
	VCS string
}

// PackageDeleteRequest defines structure for package deleting request.
type PackageDeleteRequest struct {
	// OriginalPath is a package original path without domain part.
	// E.g. for package "go.example.tld/group/pkgname" you should
	// put here "/group/pkgname".
	OriginalPath string
}

// PackageGetRequest defined structure for package information getting
// request.
type PackageGetRequest struct {
	// Should all packages be returned?
	All bool
	// If All = false, then what package name (or names) to return?
	// They should be delimited with comma in CLI.
	PackageNames []string
}
