package searchJSON

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
)

// Estructura para mapear el JSON
type Partido struct {
    Liga string `json:"liga"`
}

func SearchJSONWordKey(keyword string, folderPath string) {
    err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if filepath.Ext(path) == ".json" {
            fileContent, err := ioutil.ReadFile(path)
            if err != nil {
                return err
            }

            var partidos []Partido
            err = json.Unmarshal(fileContent, &partidos)
            if err != nil {
                log.Printf("Error al parsear el archivo JSON %s: %v", path, err)
                return nil
            }

            for _, partido := range partidos {
                if partido.Liga == keyword {
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