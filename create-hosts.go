// This script in golang will perform the following tasks
// Create a file named all_hosts.cfg with 100 hosts with the ping service check, the secuence of the names will be host###
// It will create an entry to the nagios.cfg pointing the file cfg_file=/usr/local/nagios/etc/objects/all_hosts.cfg
// Assign user and group permission to the file all_hosts.cfg
// Nagios service will be restarted

package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    const numberOfHosts = 100
    const configPath = "/usr/local/nagios/etc/objects/" // Asegúrate de cambiar esto por la ruta correcta
    const configFile = configPath + "all_hosts.cfg"

    file, err := os.Create(configFile)
    if err != nil {
        fmt.Println("Error al crear el archivo de configuración:", err)
        return
    }
    defer file.Close()

    for i := 1; i <= numberOfHosts; i++ {
        hostname := fmt.Sprintf("host%03d", i)
        ipAddress := fmt.Sprintf("192.168.1.%d", i) // Modifica esto según tu esquema de red

        configData := fmt.Sprintf(`
define host {
    use                     linux-server
    host_name               %s
    alias                   %s
    address                 %s
}
define service {
    use                     generic-service
    host_name               %s
    service_description     PING
    check_command           check_ping!100.0,20%%!500.0,60%%
}
`, hostname, hostname, ipAddress, hostname)

        if _, err := file.WriteString(configData); err != nil {
            fmt.Printf("Error al escribir en el archivo de configuración para %s: %v\n", hostname, err)
        }
    }

    // Cambiar los permisos del archivo
    if err := os.Chmod(configFile, 0644); err != nil {
        fmt.Println("Error al cambiar los permisos del archivo:", err)
    }

    // Agregar entrada en nagios.cfg
    nagiosCfg := "/usr/local/nagios/etc/nagios.cfg"
    addCfgEntry(nagiosCfg, configFile)

    // Reiniciar el servicio de Nagios
    restartNagios()
}

func addCfgEntry(nagiosCfg, configFile string) {
    cmd := exec.Command("sh", "-c", fmt.Sprintf("echo 'cfg_file=%s' >> %s", configFile, nagiosCfg))
    if err := cmd.Run(); err != nil {
        fmt.Println("Error al agregar entrada en nagios.cfg:", err)
    }
}
        if _, err := file.WriteString(configData); err != nil {
            fmt.Printf("Error al escribir en el archivo de configuración para %s: %v\n", hostname, err)
        }
    }

    // Cambiar los permisos del archivo
    if err := os.Chmod(configFile, 0644); err != nil {
        fmt.Println("Error al cambiar los permisos del archivo:", err)
    }

    // Agregar entrada en nagios.cfg
    nagiosCfg := "/usr/local/nagios/etc/nagios.cfg"
    addCfgEntry(nagiosCfg, configFile)

    // Reiniciar el servicio de Nagios
    restartNagios()
}

func addCfgEntry(nagiosCfg, configFile string) {
    cmd := exec.Command("sh", "-c", fmt.Sprintf("echo 'cfg_file=%s' >> %s", configFile, nagiosCfg))
    if err := cmd.Run(); err != nil {
        fmt.Println("Error al agregar entrada en nagios.cfg:", err)
    }
}

func restartNagios() {
    cmd := exec.Command("systemctl", "restart", "nagios")
    if err := cmd.Run(); err != nil {
        fmt.Println("Error al reiniciar el servicio de Nagios:", err)
    }
}
