package main

import (
	"fmt"
	"os/exec"
)

// check pacote instalado
func isPackageInstalled(pkgName string) bool {
	cmd := exec.Command("dpkg", "-s", pkgName)
	err := cmd.Run()
	return err == nil
}

// instala pacote
func installPackage(pkgName string) error {
	cmd := exec.Command("sudo", "apt-get", "install", "-y", pkgName)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

// inicia serviço
func startService(serviceName string) error {
	cmd := exec.Command("sudo", "systemctl", "start", serviceName)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

// reinicia serviço
func restartService(serviceName string) error {
	cmd := exec.Command("sudo", "systemctl", "restart", serviceName)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

// habilita serviço
func enableService(serviceName string) error {
	cmd := exec.Command("sudo", "systemctl", "enable", serviceName)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}


func main() {
	
	packages := []string{"apache2", "php8.0"}
	services := []string{"apache2"}

	for _, pkg := range packages {
		if isPackageInstalled(pkg) {
			fmt.Printf("O pacote %s já está instalado.\n", pkg)
		} else {
			fmt.Printf("O pacote %s não está instalado. Instalando...\n", pkg)
			err := installPackage(pkg)
			if err != nil {
				fmt.Printf("Falha ao instalar o pacote %s: %v\n", pkg, err)
			} else {
				fmt.Printf("Pacote %s instalado com sucesso.\n", pkg)
			}
		}
	}

	for _, svc := range services {
		fmt.Printf("Habilitando o serviço %s...\n", svc)
		err := enableService(svc)
		if err != nil {
			fmt.Printf("Falha ao habilitar o serviço %s: %v\n", svc, err)
		} else {
			fmt.Printf("Serviço %s habilitado com sucesso.\n", svc)
		}

		fmt.Printf("Reiniciando o serviço %s...\n", svc)
		err = restartService(svc)
		if err != nil {
			fmt.Printf("Falha ao reiniciar o serviço %s: %v\n", svc, err)
		} else {
			fmt.Printf("Serviço %s reiniciado com sucesso.\n", svc)
		}
	}
}