package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

// childCmd represents the child command
var childCmd = &cobra.Command{
	Use: "child",
	Run: func(cmd *cobra.Command, args []string) {
		syscall.Sethostname([]byte("container"))

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

		rootfsPath := cwd + "/rootfs"

		// NOTE: Change the root directory
		if err := syscall.Chroot(rootfsPath); err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

		// NOTE: Change directory to root in the new filesystem
		if err := syscall.Chdir("/"); err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

		// NOTE: Mount /proc in the new filesystem
		if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

		// Setup cgroups for resource management
		setupCgroups()

		command := exec.Command(args[0], args[1:]...)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		if err := command.Run(); err != nil {
			fmt.Println("ERROR", err)
			os.Exit(1)
		}
		// NOTE: Unmount /proc after execution
		if err := syscall.Unmount("/proc", 0); err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}
	},
}

// setupCgroups sets up the cgroups for resource management.
func setupCgroups() {
	// NOTE: Create a new cgroup directory
	cgroupPath := "/sys/fs/cgroup/my_cgroup"
	if err := os.MkdirAll(cgroupPath, 0755); err != nil {
		fmt.Println("ERROR creating cgroup directory:", err)
		os.Exit(1)
	}

	// Set resource limits (example: setting memory limit to 100MB)
	memLimit := "100M"
	memLimitPath := cgroupPath + "/memory.max"
	if err := os.WriteFile(memLimitPath, []byte(memLimit), 0644); err != nil {
		fmt.Println("ERROR setting memory limit:", err)
		os.Exit(1)
	}

	// NOTE: Add the current process to the cgroup
	cgroupProcsPath := cgroupPath + "/cgroup.procs"
	if err := os.WriteFile(cgroupProcsPath, []byte(fmt.Sprintf("%d", syscall.Getpid())), 0644); err != nil {
		fmt.Println("ERROR adding process to cgroup:", err)
		os.Exit(1)
	}
}
func init() {
	rootCmd.AddCommand(childCmd)
}
