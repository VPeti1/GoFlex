package flex

// If the GO garbage collector would be good it would collect this entire code and trow it into /bin

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

func failosr() {
	fmt.Println("Failed to run OSR funtion!")
	time.Sleep(500 * time.Millisecond)
	os.Exit(1)
}

func fail() {
	fmt.Println("Unsupported FLEXPKG file for this os!")
	time.Sleep(500 * time.Millisecond)
	os.Exit(1)
}

func osrmain(osreturn int) {
	if osreturn == 0 {
		return
	} else if osreturn == 1 {
		systemType := detectOS()
		if systemType == "windows" {
			fail()
		} else {

		}
	} else if osreturn == 7 {
		systemType := detectOS()
		if systemType == "windows" {

		} else {
			fail()
		}
	} else if osreturn == 2 {
		systemType := detectOS()
		if systemType == "windows" {
			fail()
		} else {
			if _, err := os.Stat("/usr/flex/arch.cw"); err == nil {

			} else if _, err := os.Stat("/usr/flex/deb.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/fed.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/suse.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/void.cw"); err == nil {
				fail()
			} else {
				failosr()
			}
		}
	} else if osreturn == 3 {
		systemType := detectOS()
		if systemType == "windows" {
			fail()
		} else {
			if _, err := os.Stat("/usr/flex/arch.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/deb.cw"); err == nil {

			} else if _, err := os.Stat("/usr/flex/fed.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/suse.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/void.cw"); err == nil {
				fail()
			} else {
				failosr()
			}
		}
	} else if osreturn == 4 {
		systemType := detectOS()
		if systemType == "windows" {
			fail()
		} else {
			if _, err := os.Stat("/usr/flex/arch.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/deb.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/fed.cw"); err == nil {

			} else if _, err := os.Stat("/usr/flex/suse.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/void.cw"); err == nil {
				fail()
			} else {
				failosr()
			}
		}
	} else if osreturn == 5 {
		systemType := detectOS()
		if systemType == "windows" {
			fail()
		} else {
			if _, err := os.Stat("/usr/flex/arch.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/deb.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/fed.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/suse.cw"); err == nil {

			} else if _, err := os.Stat("/usr/flex/void.cw"); err == nil {
				fail()
			} else {
				failosr()
			}
		}
	} else if osreturn == 6 {
		systemType := detectOS()
		if systemType == "windows" {
			fail()
		} else {
			if _, err := os.Stat("/usr/flex/arch.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/deb.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/fed.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/suse.cw"); err == nil {
				fail()
			} else if _, err := os.Stat("/usr/flex/void.cw"); err == nil {

			} else {
				failosr()
			}
		}
	}

}

func checkOSRLine(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "osr = ") {
			textAfterOSR := strings.TrimSpace(strings.Split(line, "osr = ")[1])
			var os int
			switch textAfterOSR {
			case "lin":
				os = 1
			case "linarch":
				os = 2
			case "lindeb":
				os = 3
			case "linfed":
				os = 4
			case "linvoid":
				os = 5
			case "linsuse":
				os = 6
			default:
				os = 7
			}
			osrmain(os)
			break
		}
	}
}

func osrf() {
	checkOSRLine("flex.pkg")
}

func batch() {
	filePath := "flex.pkg"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to read from the FLEXPKG file: %v\n", err)
		return
	}
	defer file.Close()

	var batchLine string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "batch =") {
			batchLine = strings.TrimSpace(line)
			break
		}
	}

	if batchLine != "" {
		flex := strings.TrimPrefix(batchLine, "batch = ")
		fmt.Printf("Found batch file invoke: %s\n", flex)
		time.Sleep(1500 * time.Millisecond)
		err := exec.Command("cmd", "/c", flex).Run()
		if err != nil {
			fmt.Printf("Error executing batch command: %v\n", err)
		}
	} else {
		fmt.Println("No batch line found in the file.")
	}
}

func sh() {
	filePath := "flex.pkg"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to read from the FLEXPKG file: %v\n", err)
		return
	}
	defer file.Close()

	var shLine string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "sh =") {
			shLine = strings.TrimSpace(strings.Replace(line, "sh =", "", 1))
			break
		}
	}

	if shLine != "" {
		fmt.Printf("Found sh file: %s\n", shLine)
		pm := "sudo sh " + shLine
		time.Sleep(1500 * time.Millisecond)
		cmd := exec.Command("sh", "-c", pm)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error executing the command: %v\n", err)
		}
	}
}

func flat() {
	filePath := "flex.pkg"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to read from the FLEXPKG file: %v\n", err)
		return
	}
	defer file.Close()

	var flatpakLine string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "flatpak =") {
			flatpakLine = strings.TrimSpace(line)
			break
		}
	}

	if flatpakLine != "" {
		flex := strings.TrimPrefix(flatpakLine, "flatpak = ")
		fmt.Printf("Found packages: %s\n", flex)
		pm := "flatpak install " + flex
		time.Sleep(1500 * time.Millisecond)
		cmd := exec.Command("sh", "-c", pm)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error executing flatpak install: %v\n", err)
		}
	}
}

func cmdc() {
	//I spent so mutch time debugging this funtion because of the linux env
	//The fix was that i needed to do whatever is on line 28-36
	filePath := "flex.pkg"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("File not found:", filePath)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pkgLine string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) >= 5 && line[:5] == "cmd =" {
			pkgLine = line
			break
		}
	}
	if pkgLine != "" {
		data, err := os.ReadFile("flex.pkg")
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
		scanner := bufio.NewScanner(strings.NewReader(string(data)))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "cmd = ") {
				command := strings.TrimPrefix(line, "cmd = ")
				cmd := exec.Command("/bin/sh", "-c", command)
				err := cmd.Start()
				if err != nil {
					log.Fatal(err)
				}
				err = cmd.Wait()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Command executed successfully")
			}
			if strings.HasPrefix(line, "cmd2 = ") {
				command := strings.TrimPrefix(line, "cmd2 = ")
				cmd := exec.Command("/bin/sh", "-c", command)
				err := cmd.Start()
				if err != nil {
					log.Fatal(err)
				}
				err = cmd.Wait()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Command executed successfully")
			}
			if strings.HasPrefix(line, "cmd3 = ") {
				command := strings.TrimPrefix(line, "cmd3 = ")
				cmd := exec.Command("/bin/sh", "-c", command)
				err := cmd.Start()
				if err != nil {
					log.Fatal(err)
				}
				err = cmd.Wait()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Command executed successfully")
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error scanning file: %v", err)
		}
	}
}

func wget() {
	filePath := "flex.pkg"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("File not found:", filePath)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pkgLine string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) >= 5 && line[:5] == "get =" {
			pkgLine = line
			break
		}
	}
	if pkgLine != "" {
		if _, err := os.Stat(filePath); err == nil {
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Println("Error reading from file!")
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			var get, get2, get3 string
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, "get =") {
					get = strings.Fields(strings.Split(line, "get =")[1])[0]
				}
				if strings.Contains(line, "get2 =") {
					get2 = strings.Fields(strings.Split(line, "get2 =")[1])[0]
				}
				if strings.Contains(line, "get3 =") {
					get3 = strings.Fields(strings.Split(line, "get3 =")[1])[0]
				}
			}

			fmt.Println("Downloading the files")
			if get != "" {
				fmt.Println(get)
				exec.Command("wget", get).Run()
			}
			if get2 != "" {
				fmt.Println(get2)
				exec.Command("wget", get2).Run()
			}
			if get3 != "" {
				fmt.Println(get3)
				exec.Command("wget", get3).Run()
			}

			fmt.Println("Files downloaded successfully")
		} else {
			fmt.Println("Error reading from file!")
		}
	}
}

func git() {
	filePath := "flex.pkg"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("File not found:", filePath)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pkgLine string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) >= 5 && line[:5] == "git =" {
			pkgLine = line
			break
		}
	}
	if pkgLine != "" {
		if _, err := os.Stat(filePath); err == nil {
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Println("Error opening the file!")
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			var content strings.Builder
			for scanner.Scan() {
				content.WriteString(scanner.Text() + "\n")
			}
			git := extractValue(content.String(), "git =")
			git2 := extractValue(content.String(), "git2 =")
			git3 := extractValue(content.String(), "git3 =")
			fmt.Println("Cloning these with git")
			cloneWithGit(git)
			cloneWithGit(git2)
			cloneWithGit(git3)
			fmt.Println("Repos cloned successfully")
		} else {
			fmt.Println("Error reading from file!")
		}
	}
}

func extractValue(content, key string) string {
	if strings.Contains(content, key) {
		return strings.Fields(strings.Split(content, key)[1])[0]
	}
	return ""
}

func cloneWithGit(repo string) {
	if repo != "" {
		fmt.Println(repo)
		cmd := exec.Command("git", "clone", repo)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error cloning repo: %s\n", err)
		}
	}
}

func detectOS() string {
	if runtime.GOOS == "linux" {
		return "linux"
	}
	return "windows"
}

func pkg() {
	filePath := "flex.pkg"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("File not found:", filePath)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pkgLine string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) >= 5 && line[:5] == "pkg =" {
			pkgLine = line
			break
		}
	}

	if pkgLine != "" {
		pm := ""
		flex := strings.TrimSpace(pkgLine[6:]) // Trim spaces and assume "pkg = " is at the start of the line
		fmt.Printf("Found packages: %s\n", flex)
		switch detectOS() {
		case "linux":
			if _, err := os.Stat("/usr/flex/arch.cw"); err == nil {
				pm = "sudo pacman -S "
			} else if _, err := os.Stat("/usr/flex/deb.cw"); err == nil {
				pm = "sudo apt install "
			} else if _, err := os.Stat("/usr/flex/fed.cw"); err == nil {
				pm = "sudo yum install "
			} else if _, err := os.Stat("/usr/flex/suse.cw"); err == nil {
				pm = "sudo zypper install "
			} else if _, err := os.Stat("/usr/flex/void.cw"); err == nil {
				pm = "sudo xbps-install "
			}
		default:
			pm = "choco install "
		}
		time.Sleep(1500 * time.Millisecond)
		cmd := exec.Command("sh", "-c", pm+flex)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				fmt.Printf("Installation failed with exit code: %d\n", exitError.ExitCode())
				os.Exit(1)
			} else {
				fmt.Println("Error installing packages:", err)
				os.Exit(1)
			}
			return
		}
	}

}

func shell(os string) {
	if os == "linux" {
		flat()
		sh()
	} else {
		batch()
	}
}

func osys() {
	chk := detectOS()
	if chk == "linux" {
		files := []string{"/usr/flex/arch.cw", "/usr/flex/deb.cw", "/usr/flex/fed.cw", "/usr/flex/void.cw", "/usr/flex/suse.cw"}
		for _, file := range files {
			if _, err := os.Stat(file); err == nil {
				return
			}
		}
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("What os are you using? (arch/debian/fedora/void/opensuse): \n(Derivatives included)\n")
		distro, _ := reader.ReadString('\n')
		distro = strings.TrimSpace(strings.ToLower(distro))
		switch distro {
		case "arch":
			exec.Command("sudo", "touch", "/usr/flex/arch.cw").Run()
		case "debian":
			exec.Command("sudo", "touch", "/usr/flex/deb.cw").Run()
		case "fedora":
			exec.Command("sudo", "touch", "/usr/flex/fed.cw").Run()
		case "void":
			exec.Command("sudo", "touch", "/usr/flex/void.cw").Run()
		case "opensuse":
			exec.Command("sudo", "touch", "/usr/flex/suse.cw").Run()
		default:
			fmt.Println("Unknown os. Please enter 'arch', 'debian', 'void', 'opensuse' or 'fedora'.")
			time.Sleep(500 * time.Millisecond)
			exec.Command("clear").Run()
			osys()
		}
	} else {
		_, err := os.Stat(`C:\ProgramData\chocolatey\bin\choco.exe`)
		if os.IsNotExist(err) {
			fmt.Println("Chocolatey not installed!.")
			os.Exit(1)
		}
		time.Sleep(500 * time.Millisecond)

	}
}

func main() {
	debug.SetGCPercent(-1)
	fmt.Println(" /$$$$$$$$ /$$                     /$$$$$$$  /$$                \n| $$_____/| $$                    | $$__  $$| $$                \n| $$      | $$  /$$$$$$  /$$   /$$| $$  \\ $$| $$   /$$  /$$$$$$ \n| $$$$$   | $$ /$$__  $$|  $$ /$$/| $$$$$$$/| $$  /$$/ /$$__  $$\n| $$__/   | $$| $$$$$$$$ \\  $$$$/ | $$____/ | $$$$$$/ | $$  \\ $$\n| $$      | $$| $$_____/  >$$  $$ | $$      | $$_  $$ | $$  | $$\n| $$      | $$|  $$$$$$$ /$$/\\  $$| $$      | $$ \\  $$|  $$$$$$$\n|__/      |__/ \\_______/|__/  \\__/|__/      |__/  \\__/ \\____  $$\n                                                       /$$  \\ $$\n                                                      |  $$$$$$/\n                                                       \\______/ ")
	time.Sleep(1 * time.Second)
	filePath := "flex.pkg"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("flex.pkg file not found in current directory!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		if firstLine == "## FLEXPKG Format Version 1 By VPeti" {
			fmt.Println("Valid FLEXPKG file detected")
			time.Sleep(1 * time.Second)
			osrf()
			pkg()
			osys()
			git()
			wget()
			cmdc()
			shell(detectOS())
		} else {
			fmt.Println("The file isn't a valid FLEXPKG file!")
		}
	}
}
