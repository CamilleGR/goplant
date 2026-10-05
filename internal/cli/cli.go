package cli

import (
	"fmt"
	"goplant/internal/builder"
	"goplant/internal/server"
	"goplant/internal/template"
	"math/rand"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var plants = []string{
	"Rose", "Sunflower", "Cactus", "Maple", "Oak", "Bamboo", "Aloe", "Tulip",
	"Hibiscus", "Sequoia", "Lavender", "Mint", "Venusflytrap", "Baobab", "Lotus",
	"Orchid", "Daisy", "Fern", "Bonsai", "Poinsettia",
}

func generatePlantName(targetOS string) string {
	randomPlant := plants[rand.Intn(len(plants))]
	ext := getExtension(targetOS)
	return strings.ToLower(randomPlant) + ext
}

func getExtension(targetOS string) string {
	if targetOS == "windows" {
		return ".exe"
	}
	return ""
}

// RootCmd est la commande racine de l'application
var RootCmd = &cobra.Command{
	Use:   "goplant",
	Short: "Generate simple binary reverse shell implant",
	Long:  "Goplant generates executable reverse shell implants for pentesting",
}

// GenerateCmd est la commande pour générer un implant
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a reverse shell implant",
	Long:  "Generate an executable reverse shell implant binary for the specified listener",
	Run: func(cmd *cobra.Command, args []string) {
		// Récupérer les flags
		lhost, _ := cmd.Flags().GetString("lhost")
		lport, _ := cmd.Flags().GetString("lport")
		output, _ := cmd.Flags().GetString("output")
		templateName, _ := cmd.Flags().GetString("template")
		targetOS, _ := cmd.Flags().GetString("os")
		arch, _ := cmd.Flags().GetString("arch")

		// Si output n'est pas défini, générer un nom de plante
		if output == "" {
			output = generatePlantName(targetOS)
		}

		// Récupérer le template
		tmpl, err := template.GetTemplate(templateName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Template utilisé: %s\n", templateName)
		fmt.Printf("Cible: %s/%s\n", targetOS, arch)
		fmt.Printf("Output: %s\n", output)

		// Utiliser le builder pour générer l'implant
		if err := builder.Build(tmpl, lhost, lport, output, targetOS, arch); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

// ServeCmd est la commande pour servir des implants
var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start an implant server",
	Long:  "Start a server to serve generated implants",
	Run: func(cmd *cobra.Command, args []string) {
		// Récupérer les flags
		lhost, _ := cmd.Flags().GetString("lhost")
		lport, _ := cmd.Flags().GetString("lport")
		bindAddr, _ := cmd.Flags().GetString("bindHost")
		bindPort, _ := cmd.Flags().GetString("bindPort")
		targetOS, _ := cmd.Flags().GetString("os")
		arch, _ := cmd.Flags().GetString("arch")

		// Créer et démarrer le serveur
		srv := server.New(lhost, lport, bindAddr, bindPort, targetOS, arch)
		if err := srv.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
			os.Exit(1)
		}
	},
}

// ListCmd est la commande pour lister les templates disponibles
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available templates",
	Long:  "List all available implant templates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available templates:")
		for _, tmpl := range template.GetAvailableTemplates() {
			fmt.Printf("- [%s] %s\n", tmpl.GetName(), tmpl.GetDesc())
		}
	},
}

func init() {
	// Ajouter les sous-commandes
	RootCmd.AddCommand(GenerateCmd)
	RootCmd.AddCommand(ServeCmd)
	RootCmd.AddCommand(ListCmd)

	// Flags pour generate
	GenerateCmd.Flags().StringP("lhost", "", "", "Listener host (required)")
	GenerateCmd.Flags().StringP("lport", "", "", "Listener port (required)")
	GenerateCmd.Flags().StringP("output", "o", "", "Output file path (optional, random plant name if not specified)")
	GenerateCmd.Flags().StringP("template", "t", "go", "Implant template to use")
	GenerateCmd.Flags().StringP("os", "", "", "Target OS (required)")
	GenerateCmd.Flags().StringP("arch", "a", "", "Target architecture [arm64, amd64, x86](required)")
	GenerateCmd.MarkFlagRequired("lhost")
	GenerateCmd.MarkFlagRequired("lport")
	GenerateCmd.MarkFlagRequired("os")
	GenerateCmd.MarkFlagRequired("a")

	// Flags pour serve
	ServeCmd.Flags().StringP("lhost", "", "", "Listener host (required)")
	ServeCmd.Flags().StringP("lport", "", "", "Listener port (required)")
	ServeCmd.Flags().StringP("template", "t", "go", "Implant template to use")
	ServeCmd.Flags().StringP("bindHost", "H", "0.0.0.0", "Bind address for the server")
	ServeCmd.Flags().StringP("bindPort", "p", "1337", "Bind port for the server")
	ServeCmd.Flags().StringP("os", "", "", "Target OS (required)")
	ServeCmd.Flags().StringP("arch", "a", "", "Target architecture (required)")
	ServeCmd.MarkFlagRequired("lhost")
	ServeCmd.MarkFlagRequired("lport")
	ServeCmd.MarkFlagRequired("os")
	ServeCmd.MarkFlagRequired("a")
}
