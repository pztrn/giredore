package structs

const (
	ErrInvalidAllowedIPDefined             Error = "Invalid allowed IP address defined."
	ErrIPAddressNotAllowed                 Error = "IP address not allowed to access configuration API."
	ErrNoPackagesFound                     Error = "No packages found."
	ErrPackageOrigPathShouldStartWithSlash Error = "Package's original name (path) should start with slash."
	ErrPackageWasntDefined                 Error = "Passed package wasn't defined."
	ErrParsingAllowedIPsSetRequest         Error = "Error parsing allowed IPs request."
	ErrParsingDeleteRequest                Error = "Delete request parsing failed"
	ErrParsingPackagesGetRequest           Error = "Error parsing package(s) info get request"
)

type Error string
