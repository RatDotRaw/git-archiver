package sysinfo

import (
	"fmt"
	"os"
	// "path/filepath"
	"runtime"
	"strings"
	"time"
)

type SystemInfo struct {
	// Basic information
	Hostname     string
	Runtime      string
	Architecture string

	// CPU information
	NumCPU     int
	GoMaxProcs int

	// Memory information
	TotalMemory uint64 // This will be populated differently based on OS

	// OS specific
	OSVersion string

	// Runtime information
	GoVersion string

	// Time information
	Timezone    string
	CurrentTime string

	// User information
	Username string
	HomeDir  string

	// File system
	WorkingDir string
	TempDir    string

	// Environment
	EnvVars map[string]string
}




// GetSystemInfo gathers comprehensive system information
func GetSystemInfo() SystemInfo {
	var info SystemInfo

	// Basic system info
	hostname, err := os.Hostname()
	if err == nil {
		info.Hostname = hostname
	} else {
		fmt.Errorf("Error getting hostname: %s\n", err)
	}

	info.Runtime = runtime.GOOS
	info.Architecture = runtime.GOARCH

	// CPU information
	info.NumCPU = runtime.NumCPU()
	info.GoMaxProcs = runtime.GOMAXPROCS(0) // Gets current value without changing it

	// OS Version - platform specific implementations would go here
	// This is a simplified version
	info.OSVersion = getOSVersion()

	// Go runtime information
	info.GoVersion = runtime.Version()

	// Time information
	localZone, _ := time.Now().Zone()
	info.Timezone = localZone
	info.CurrentTime = time.Now().Format(time.RFC3339)

	// User information
	info.Username = os.Getenv("USER") // On Windows, use "USERNAME"
	info.HomeDir, _ = os.UserHomeDir()

	// File system
	info.WorkingDir, _ = os.Getwd()
	info.TempDir = os.TempDir()

	// Environment variables
	info.EnvVars = make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			info.EnvVars[pair[0]] = pair[1]
		}
	}

	// Memory information would typically require platform-specific code or external libraries
	// For a simple cross-platform solution, we'll leave it as 0 for now
	info.TotalMemory = getMemoryInfo()

	return info
}

// getOSVersion attempts to get OS version information
// This is a simplified function - actual implementation would vary by OS
func getOSVersion() string {
	switch runtime.GOOS {
	case "windows":
		return os.Getenv("OS")
	case "darwin":
		// On macOS, reading from sw_vers would be ideal
		return "macOS (version retrievable via 'sw_vers')"
	case "linux":
		// On Linux, reading from /etc/os-release or similar would be ideal
		if _, err := os.Stat("/etc/os-release"); err == nil {
			content, err := os.ReadFile("/etc/os-release")
			if err == nil {
				return extractLinuxVersion(string(content))
			}
		}
		return "Linux (version information unavailable)"
	default:
		return "Unknown OS version"
	}
}

// extractLinuxVersion parses /etc/os-release content
func extractLinuxVersion(content string) string {
	lines := strings.Split(content, "\n")
	prettyName := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			prettyName = strings.Trim(line[12:], "\"")
			break
		}
	}
	if prettyName != "" {
		return prettyName
	}
	return "Linux"
}

// getMemoryInfo is a placeholder for actual memory information retrieval
// Actual implementation would use OS-specific methods or libraries
func getMemoryInfo() uint64 {
	// To implement this properly, use:
	// - Windows: GlobalMemoryStatusEx or similar WinAPI call
	// - Linux: Read from /proc/meminfo
	// - macOS: Use syscall or host_statistics API
	return 0 // Placeholder
}

// PrintSystemInfo displays system information
func PrettyPrintSystemInfo(info SystemInfo) {
	fmt.Println("=== System Information ===")
	fmt.Printf("Hostname: %s\n", info.Hostname)
	fmt.Printf("OS: %s (%s)\n", info.Runtime, info.OSVersion)
	fmt.Printf("Architecture: %s\n", info.Architecture)
	fmt.Printf("CPUs: %d (GOMAXPROCS: %d)\n", info.NumCPU, info.GoMaxProcs)
	fmt.Printf("Go Version: %s\n", info.GoVersion)
	fmt.Printf("Time: %s (%s)\n", info.CurrentTime, info.Timezone)
	fmt.Printf("User: %s\n", info.Username)
	fmt.Printf("Home Directory: %s\n", info.HomeDir)
	fmt.Printf("Working Directory: %s\n", info.WorkingDir)
	fmt.Printf("Temp Directory: %s\n", info.TempDir)

	if info.TotalMemory > 0 {
		fmt.Printf("Total Memory: %d bytes\n", info.TotalMemory)
	}

	fmt.Println("\n=== Environment Variables ===")
	// Print only a few important environment variables
	envVarsToShow := []string{"PATH", "GOPATH", "GOROOT", "HOME", "USER"}
	for _, key := range envVarsToShow {
		if val, ok := info.EnvVars[key]; ok {
			fmt.Printf("%s: %s\n", key, val)
		}
	}
	fmt.Println("==========================")
}
