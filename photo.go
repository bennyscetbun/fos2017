package main

import (
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"path/filepath"

	"encoding/hex"

	"github.com/nfnt/resize"
)

func uploadPhoto(w http.ResponseWriter, r *http.Request) {
	userInter, err := ab.CurrentUser(w, r)
	if internalError(w, err) {
		return
	}
	if userInter == nil {
		badRequest(w, fmt.Errorf("Not logged"))
		return
	}
	user := userInter.(*User)
	r.ParseMultipartForm(32 << 20)
	file, _ /*handler*/, err := r.FormFile("uploadfile")
	if internalError(w, err) {
		return
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err == image.ErrFormat {
		badRequest(w, err)
		return
	} else if internalError(w, err) {
		return
	}
	m := resize.Resize(1000, 0, img, resize.Lanczos3)
	photoName := fmt.Sprintf("%s.jpg", hex.EncodeToString([]byte(user.ID)))
	filep := filepath.Join(os.Getenv("PHOTO_PATH"), photoName)
	f, err := os.OpenFile(filep, os.O_WRONLY|os.O_CREATE, 0666)
	if internalError(w, err) {
		return
	}
	if internalError(w, jpeg.Encode(f, m, &jpeg.Options{Quality: 80})) {
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, fmt.Sprintf("/assets/%s", photoName))
}
