package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"go.dev.pztrn.name/giredore/internal/structs"
)

// This structure represents configuration that will be parsed via file.
// Despite on exporting fields there are setters and getters defined because
// data from configuration file will be parsed in exported fields and they
// may be accesses concurrently. In other words DO NOT USE EXPORTED FIELDS
// DIRECTLY!
type fileConfig struct {
	packagesMutex sync.RWMutex
	// Packages describes packages mapping.
	Packages map[string]*structs.Package
	// HTTP describes HTTP server configuration.
	HTTP struct {
		allowedipsmutex sync.RWMutex
		// Listen is an address on which HTTP server will listen.
		Listen string
		// AllowedIPs is a list of IPs that allowed to access API.
		// There might be other authentication implemented in future.
		AllowedIPs []string
		// WaitForSeconds is a timeout during which we will wait for
		// HTTP server be up. If timeout will pass and HTTP server won't
		// start processing requests - giredore will exit.
		WaitForSeconds int
	}
}

func (fc *fileConfig) AddOrUpdatePackage(pkg *structs.Package) {
	fc.packagesMutex.Lock()
	fc.Packages[pkg.OriginalPath] = pkg
	fc.packagesMutex.Unlock()
}

func (fc *fileConfig) DeletePackage(req *structs.PackageDeleteRequest) []structs.Error {
	var errors []structs.Error

	fc.packagesMutex.Lock()
	defer fc.packagesMutex.Unlock()
	_, found := fc.Packages[req.OriginalPath]

	if !found {
		errors = append(errors, structs.ErrPackageWasntDefined)

		return errors
	}

	delete(fc.Packages, req.OriginalPath)

	return errors
}

func (fc *fileConfig) GetAllowedIPs() []string {
	var allowedIPs []string

	fc.HTTP.allowedipsmutex.RLock()
	allowedIPs = append(allowedIPs, fc.HTTP.AllowedIPs...)
	fc.HTTP.allowedipsmutex.RUnlock()

	return allowedIPs
}

func (fc *fileConfig) GetAllPackagesInfo() map[string]*structs.Package {
	pkgs := make(map[string]*structs.Package)

	fc.packagesMutex.Lock()
	for name, pkg := range fc.Packages {
		pkgs[name] = pkg
	}
	fc.packagesMutex.Unlock()

	return pkgs
}

func (fc *fileConfig) GetPackagesInfo(packages []string) (map[string]*structs.Package, []structs.Error) {
	pkgs := make(map[string]*structs.Package)

	var errors []structs.Error

	fc.packagesMutex.Lock()
	for _, neededPkg := range packages {
		pkgData, found := fc.Packages[neededPkg]
		if !found {
			errors = append(errors, structs.ErrPackageWasntDefined+structs.Error(" Package was: "+neededPkg))
		} else {
			pkgs[neededPkg] = pkgData
		}
	}
	fc.packagesMutex.Unlock()

	return pkgs, errors
}

// Initialize parses file contents into structure.
func (fc *fileConfig) Initialize() {
	configPath := filepath.Join(envCfg.DataDir, "config.json")
	cfgLoadLog := log.With().Str("configuration path", configPath).Logger()
	cfgLoadLog.Info().Msg("Loading configuration file...")

	configPath, err := fc.normalizePath(configPath)
	if err != nil {
		cfgLoadLog.Fatal().Err(err).Msg("Failed to normalize configuration file path.")
	}

	// Check if file "config.json" specified in envConfig.DataDir field
	// exists.
	if _, err2 := os.Stat(configPath); os.IsNotExist(err2) {
		cfgLoadLog.Error().Msg("Unable to load configuration from filesystem.")

		return
	}

	// Try to load file into memory.
	fileData, err3 := ioutil.ReadFile(configPath)
	if err3 != nil {
		cfgLoadLog.Fatal().Err(err3).Msg("Failed to read configuration file data into memory.")
	}

	// ...and parse it.
	err4 := json.Unmarshal(fileData, fc)
	if err4 != nil {
		cfgLoadLog.Fatal().Err(err4).Msg("Failed to parse configuration file.")
	}

	if fc.Packages == nil {
		fc.Packages = make(map[string]*structs.Package)
	}

	// Ensure that localhost (127.0.0.1) are defined in AllowedIPs.
	var localhostIsAllowed bool

	for _, ip := range fc.HTTP.AllowedIPs {
		if strings.Contains(ip, "127.0.0.1") {
			localhostIsAllowed = true

			break
		}
	}

	if !localhostIsAllowed {
		cfgLoadLog.Warn().Msg("Localhost (127.0.0.1) wasn't allowed to access configuration API, adding it to list of allowed IP addresses")

		fc.HTTP.AllowedIPs = append(fc.HTTP.AllowedIPs, "127.0.0.1")
	} else {
		cfgLoadLog.Debug().Msg("Localhost (127.0.0.1) is allowed to access configuration API")
	}

	cfgLoadLog.Debug().Msgf("Configuration parsed: %+v", fc)
	cfgLoadLog.Info().Int("packages count", len(fc.Packages)).Msg("Packages list loaded")
}

// Normalizes passed configuration file path.
func (fc *fileConfig) normalizePath(configPath string) (string, error) {
	// Normalize configuration file path.
	if strings.Contains(configPath, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			// nolint:wrapcheck
			return "", err
		}

		configPath = strings.Replace(configPath, "~", homeDir, 1)
	}

	absPath, err1 := filepath.Abs(configPath)
	if err1 != nil {
		// nolint:wrapcheck
		return "", err1
	}

	return absPath, nil
}

// Save saves configuration into file.
func (fc *fileConfig) Save() {
	configPath := filepath.Join(envCfg.DataDir, "config.json")
	cfgSaveLog := log.With().Str("configuration path", configPath).Logger()
	cfgSaveLog.Info().Msg("Saving configuration file...")

	data, err := json.Marshal(fc)
	if err != nil {
		cfgSaveLog.Fatal().Err(err).Msg("Failed to encode data into JSON. Configuration file won't be saved!")

		return
	}

	configPath, err1 := fc.normalizePath(configPath)
	if err1 != nil {
		cfgSaveLog.Fatal().Err(err1).Msg("Failed to normalize configuration file path.")
	}

	err2 := ioutil.WriteFile(configPath, data, os.ModePerm)
	if err2 != nil {
		cfgSaveLog.Fatal().Err(err2).Msg("Failed to write configuration file data to file.")
	}

	cfgSaveLog.Info().Msg("Configuration file saved.")
}

func (fc *fileConfig) SetAllowedIPs(allowedIPs []string) {
	fc.HTTP.allowedipsmutex.Lock()
	fc.HTTP.AllowedIPs = allowedIPs
	fc.HTTP.allowedipsmutex.Unlock()
}
