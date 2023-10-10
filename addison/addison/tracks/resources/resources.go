package resources

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tracks/repository"
)

func updateTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Println("INFO updateTrack start: ", id)
	var track repository.Track
	if err := json.NewDecoder(r.Body).Decode(&track); err == nil {
		if id == track.Id {
			if n := repository.Update(track); n > 0 {
				log.Println("INFO updateTrack 204 updated: ", id, len(track.Audio))
				w.WriteHeader(204) /* No Content */
			} else if n := repository.Insert(track); n > 0 {
				log.Println("INFO updateTrack 201 created: ", id, len(track.Audio))
				w.WriteHeader(201) /* Created */
			} else {
				log.Println("ERROR updateTrack 500 failed to create/update: ", id)
				w.WriteHeader(500) /* Internal Server Error */
			}
		} else {
			log.Println("ERROR updateTrack 400 id mismatch: ", id)
			w.WriteHeader(400) /* Bad Request */
		}
	} else {
		log.Println("ERROR updateTrack 400 cannot decode body: ", id)
		w.WriteHeader(400) /* Bad Request */
	}
	log.Println("INFO updateTrack end: ", id)
}

func readTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Println("INFO readTrack start: ", id)
	if track, n := repository.Read(id); n > 0 {
		d := repository.Track{Id: track.Id, Audio: track.Audio}
		log.Println("INFO readTrack 200 found: ", id, track.Id, len(track.Audio))
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(d)
	} else if n == 0 {
		log.Println("ERROR readTrack 404 not foundi: ", id)
		w.WriteHeader(404) /* Not Found */
	} else {
		log.Println("ERROR readTrack 500 failed to read: ", id)
		w.WriteHeader(500) /* Internal Server Error */
	}
	log.Println("INFO readTrack end: ", id)
}

func listTracks(w http.ResponseWriter, r *http.Request) {
	if list, success := repository.List(); success == true {
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(list)
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func deleteTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if n := repository.Delete(id); n > 0 {
		w.WriteHeader(200) /* OK */
	} else if n == 0 {
		w.WriteHeader(404) /* Not Found */
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* Update */
	r.HandleFunc("/tracks/{id}", updateTrack).Methods("PUT")
	/* Read */
	r.HandleFunc("/tracks/{id}", readTrack).Methods("GET")
	/* List */
	r.HandleFunc("/tracks", listTracks).Methods("GET")
	/* Delete */
	r.HandleFunc("/tracks/{id}", deleteTrack).Methods("DELETE")
	return r
}
