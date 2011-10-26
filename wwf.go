package main

import csv "csv"
import fmt "fmt"
import log "log"
import os "os"
import strconv "strconv"

func main() {
    f, err := os.Open("letters.txt")
    if err != nil {
        log.Print(err)
        os.Exit(-1)
    }
    
    reader := csv.NewReader(f)
    records, err := reader.ReadAll()
    if (err != nil) {
        log.Print(err)
        os.Exit(-1)
    }
   
    var ttl_letters, ttl_pts int
    for _, row := range records {
        freq, _ := strconv.Atoi(row[1])
        pts, _ := strconv.Atoi(row[2])
        ttl_letters += freq
        ttl_pts += (pts * freq)
        fmt.Printf("Letter %s has frequency %d and value %d\n", row[0], freq, pts)
    }
    fmt.Printf("Total letters: %d. Total points: %d\n", ttl_letters, ttl_pts)
}
    
