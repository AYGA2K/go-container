package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		// Run the current executable with "child" as the argument
		command := exec.Command("/proc/self/exe", append([]string{"child"}, args...)...)

		// Redirect stdin, stdout, and stderr to the parent
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		// Set process attributes to use new namespaces for isolation
		command.SysProcAttr = &syscall.SysProcAttr{
			// NOTE: Cloneflags controls which namespaces the child process will have.
			// CLONE_NEWUTS - Creates a new UTS namespace (allows container to have a different hostname).
			// CLONE_NEWPID - Creates a new PID namespace (isolates process IDs).
			// CLONE_NEWNS - Creates a new mount namespace (isolates file system mount points).
			Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,

			// // NOTE: Unshareflags specifies what namespaces should be unshared from the parent process.
			// // By unsharing the mount namespace (CLONE_NEWNS), changes made to mounts in the container will not affect the host system.
			Unshareflags: syscall.CLONE_NEWNS,
		}

		// Execute the command, starting the container
		if err := command.Run(); err != nil {
			fmt.Println("ERROR", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
