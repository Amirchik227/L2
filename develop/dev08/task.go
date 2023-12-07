package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	go_ps "github.com/mitchellh/go-ps"
)

// cd меняет рабочую директорию
func cd(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("Error: Bad directory")
	}
}

// new_process запускает новый процесс, с соответствующими аргументами
func new_process(arg string) {
	args := strings.Split(arg, " ")

	if len(args) < 1 {
		fmt.Println("Error: Bad arguments")
	} else {
		cmd := exec.Command(args[0], args[1:]...)
		go func() {
			_ = cmd.Run()
		}()
	}
}

// ps выводит список процессов в shell
func ps() {
	pcs, err := go_ps.Processes()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("№ PID PPID EXECUTABLE")
	for i, proc := range pcs {
		fmt.Println(i, proc.Pid(), proc.PPid(), proc.Executable())
	}
}

// echo выводит аргумент в shell
func echo(arg string) {
	fmt.Println(arg)
}

// pwd выводит рабочую директорию в shell
func pwd() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(path)
	}
}

// shell запускает бесконечный цикл обработки
func shell() {
	sc := bufio.NewScanner(os.Stdin)
	for fmt.Print(">"); sc.Scan(); fmt.Print(">") {
		cmd := sc.Text()
		if cmd == "quit" {
			break
		} else {
			сommand_processing(cmd)
		}
	}
}

// kill завершает процесс по PID
func kill(pidStr string) {
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	prc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	err = prc.Kill()
	if err != nil {
		fmt.Println("Error: cannot kill a process \n", err)
	}
}

// сommand_processing обрабатывает команду, введенную в shell
func сommand_processing(cmd string) {
	switch strings.Split(cmd, " ")[0] {
	case "echo":
		echo(strings.Replace(cmd, "echo ", "", 1))
	case "pwd":
		pwd()
	case "cd":
		cd(strings.Replace(cmd, "cd ", "", 1))
	case "fork-exec":
		new_process(strings.Replace(cmd, "fork-exec ", "", 1))
	case "ps":
		ps()
	case "kill":
		kill(strings.Replace(cmd, "kill ", "", 1))
	default:
		fmt.Println("Error: Unknown command")
	}
}

func main() {
	shell()
}
