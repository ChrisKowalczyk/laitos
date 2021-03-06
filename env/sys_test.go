package env

import (
	"runtime"
	"strings"
	"testing"
)

func TestGetProgramMemUsageKB(t *testing.T) {
	if runtime.GOOS != "linux" {
		// Just make sure the function does not crash
		GetProgramMemoryUsageKB()
		return
	}
	if usage := GetProgramMemoryUsageKB(); usage < 1000 {
		t.Fatal(usage)
	}
}

func TestGetSystemMemoryUsageKB(t *testing.T) {
	if runtime.GOOS != "linux" {
		// Just make sure the function does not crash
		GetSystemMemoryUsageKB()
		return
	}
	used, total := GetSystemMemoryUsageKB()
	if used < 1000 || total < used {
		t.Fatal(used, total)
	}
}

func TestGetSystemLoad(t *testing.T) {
	if runtime.GOOS != "linux" {
		// Just make sure the function does not crash
		GetSystemMemoryUsageKB()
		return
	}
	load := GetSystemLoad()
	if len(load) < 6 {
		t.Fatal(load)
	}
}

func TestGetSystemUptimeSec(t *testing.T) {
	if runtime.GOOS != "linux" {
		// Just make sure the function does not crash
		GetSystemUptimeSec()
		return
	}
	uptime := GetSystemUptimeSec()
	if uptime < 10 {
		t.Fatal(uptime)
	}
}

func TestGetSysctl(t *testing.T) {
	key := "vm.max_map_count"
	if runtime.GOOS != "linux" {
		// Just make sure the function does not crash
		GetSysctlInt(key)
		GetSysctlStr(key)
		return
	}
	if val, err := GetSysctlStr(key); err != nil || val == "" {
		t.Fatal(val, err)
	}
	if val, err := GetSysctlInt(key); err != nil || val < 1 {
		t.Fatal(val, err)
	}
	if old, err := IncreaseSysctlInt(key, 65535); old == 0 ||
		(err != nil && !strings.Contains(err.Error(), "permission") && !strings.Contains(err.Error(), "read-only")) {
		t.Fatal(err)
	}
}
