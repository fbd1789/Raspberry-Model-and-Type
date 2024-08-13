package main

import (
	"bytes"
	"fmt"
	"log"
    "strings"
	// "os"
	"flag"
	"golang.org/x/crypto/ssh"
	// "gopkg.in/ini.v1"
)


// SystemInfo represente les informations generales du système.
type SystemInfo struct {
	Revision  string
	Serial    string
	Model     string
	Mem		  string
}

func main() {
	// Fichier config.ini
	// inidata, err := ini.Load("config.ini")
	// if err != nil {
	//    fmt.Printf("Fail to read file: %v", err)
	//    os.Exit(1)
	//  }
	// UserName := inidata.Section("global").Key("username").String()
	// Password := inidata.Section("global").Key("password").String()
	// Host := inidata.Section("global").Key("device").String()
	// ------------------------------------------------------------

	// Sans fichier config.ini mais avec des flags
	// Définir les flags pour l'entrée utilisateur via la CLI
	user := flag.String("user", "", "Nom d'utilisateur SSH")
	password := flag.String("password", "", "Mot de passe SSH")
	host := flag.String("host", "", "Adresse IP ou hôte du serveur SSH")

	// Parse les flags
	flag.Parse()

	// Vérifier si les flags sont fournis
	if *user == "" || *password == "" || *host == "" {
		log.Fatalf("Veuillez fournir les paramètres user, password et host")
	}
	UserName := *user
	Password := *password
	Host := *host


	// Configuration SSH
	config := &ssh.ClientConfig{
		User: UserName,
		Auth: []ssh.AuthMethod{
			ssh.Password(Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Pour ignorer la verification de l'empreinte du serveur
	}

	// Connexion SSH
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:22",Host), config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer client.Close()

	// Creation de la session
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	// Recuperation de la sortie de la commande
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	// Execution des commandes pour le modele et la memoire
	cmd := "cat /proc/cpuinfo; free"
	if err := session.Run(cmd); err != nil {
		log.Fatalf("Failed to run: %s", err)
	}

	// Traitement du resultat
	// Conversion de l'hexa en string
	output := stdoutBuf.String()
	// Delimitation des retours de ligne
	lines := strings.Split(output, "\n")

	// Definition pour la structure
	var systemInfo SystemInfo
	// Affichage des lignes delimitees
	for _, line := range lines {
		// Enlever les espaces en debut et fin de chaine
		line = strings.TrimSpace(line)
		// Separation de la cle et de la valeur
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		// Stockage cle:valeur
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remplir les informations du système et de la memoire en fonction de la cle
		switch key {
		case "Revision":
			systemInfo.Revision = value
		case "Serial":
			systemInfo.Serial = value
		case "Model":
			systemInfo.Model = value
		case "Mem":
			systemInfo.Mem = strings.Split(value, " ")[0]  //Il faut prendre uniquement la premiere valeur
		}
	}
	fmt.Printf("Revision: %s\n",systemInfo.Revision)
	fmt.Printf("Serial: %s\n",systemInfo.Serial)
	fmt.Printf("Model: %s\n",systemInfo.Model)
	fmt.Printf("Mem: %s\n",systemInfo.Mem)
}
