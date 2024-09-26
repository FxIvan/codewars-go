package searchJSON

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// Estructuras para mapear el JSON
type Partido struct {
    Liga string `json:"nombre"`
}

type TeamGroup struct {
    Id            string `json:"id"`
    Nombre        string `json:"nombre"`
    NombreCorto   string `json:"nombreCorto"`
    NombreOficial string `json:"nombreOficial"`
    Codigo        string `json:"codigo"`
    Ranking       int    `json:"ranking"`
    Logo          string `json:"logo"`
    Puntos        int    `json:"puntos"`
}

type Team struct {
    Id      string     `json:"id"`
    Nombre  string     `json:"nombre"`
    Equipos []TeamGroup `json:"equipos"`
}

type GroupStage struct {
    Id     string  `json:"id"`
    Nombre string  `json:"nombre"`
    Grupos []Team  `json:"grupos"`
}

// Función para buscar una palabra clave en el campo "Liga" de archivos JSON (case-insensitive)
func SearchJSONWordKey(keyword string, folderPath string) {
    // Convertimos la palabra clave a minúsculas para hacer la búsqueda insensible a mayúsculas
    keywordLower := strings.ToLower(keyword)

    err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if filepath.Ext(path) == ".json" {
            fileContent, err := ioutil.ReadFile(path)
            if err != nil {
                return fmt.Errorf("error al leer el archivo %s: %v", path, err)
            }

            var partidos []Partido
            err = json.Unmarshal(fileContent, &partidos)
            if err != nil {
                log.Printf("Error al parsear el archivo JSON %s: %v", path, err)
                return nil
            }

            for _, partido := range partidos {
                // Convertimos el campo Liga a minúsculas para comparar de manera insensible a mayúsculas
                if strings.ToLower(partido.Liga) == keywordLower {
                    fmt.Printf("La palabra clave '%s' se encuentra en el archivo: %s\n", keyword, path)
                    break
                }
            }
        }
        return nil
    })

    if err != nil {
        log.Fatalf("Error al buscar archivos JSON: %v", err)
    }
}

// Función para buscar un string dentro de un slice (insensible a mayúsculas)
func stringInSlice(a string, list []string) bool {
    aLower := strings.ToLower(a)
    for _, b := range list {
        if strings.ToLower(b) == aLower {
            return true
        }
    }
    return false
}

// Función para buscar un equipo en los archivos JSON (case-insensitive)
func SearchObjectTeam(keyword string, folderPath string) {
    // Convertimos la palabra clave a minúsculas para hacer la búsqueda insensible a mayúsculas
    keywordLower := strings.ToLower(keyword)

    err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if filepath.Ext(path) == ".json" {
            fileContent, err := ioutil.ReadFile(path)
            if err != nil {
                return fmt.Errorf("error al leer el archivo %s: %v", path, err)
            }

            var partidos []GroupStage
            err = json.Unmarshal(fileContent, &partidos)
            if err != nil {
                log.Printf("Error al parsear el archivo JSON %s: %v", path, err)
                return nil
            }

            for _, partido := range partidos {
                for _, grupo := range partido.Grupos {
                    for _, equipo := range grupo.Equipos {
                        // Dividimos el nombre del equipo en palabras
                        equipoTeamSplit := strings.Split(equipo.Nombre, " ")

                        // Buscamos si la palabra clave está en el slice de palabras del nombre del equipo
                        if stringInSlice(keywordLower, equipoTeamSplit) {
                            fmt.Printf("El equipo '%s' se encuentra en el archivo: %s\n", keyword, path)
                            return nil // Detenemos la búsqueda cuando encontramos el equipo
                        }
                    }
                }
            }
        }
        return nil
    })

    if err != nil {
        log.Fatalf("Error al buscar archivos JSON: %v", err)
    }
}
