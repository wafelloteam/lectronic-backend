package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/wafellofazztrack/lectronic-backend/lib"
)

func AuthUploadImage() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(w, r.Body, 2*1024*1024)
			file, fileHeader, err := r.FormFile("image")
			if err != nil {
				if err == http.ErrMissingFile {
					imageName := "default_image.jpg"
					ctx := context.WithValue(r.Context(), "image", imageName)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				lib.NewRes(err.Error(), 401, true).Send(w)
				return
			}

			defer file.Close()

			buff := make([]byte, 512)
			_, err = file.Read(buff)
			if err != nil {
				lib.NewRes(err.Error(), 500, true).Send(w)
				return
			}

			filetype := http.DetectContentType(buff)
			if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" && filetype != "image/webp" {
				lib.NewRes("file format is not allowed. Please upload a JPEG, JPG or PNG image", 401, true).Send(w)
				return
			}

			_, err = file.Seek(0, io.SeekStart)
			if err != nil {
				lib.NewRes(err.Error(), 500, true).Send(w)
				return
			}

			err = os.MkdirAll("./public/image", os.ModePerm)
			if err != nil {
				lib.NewRes(err.Error(), 401, true).Send(w)
				return
			}

			imageName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
			pathRes := filepath.Join("./public/image", imageName)
			dst, err := os.Create(pathRes)
			if err != nil {
				_ = os.Remove(pathRes)
				lib.NewRes(err.Error(), 401, true).Send(w)
				return
			}

			defer dst.Close()

			_, err = io.Copy(dst, file)
			if err != nil {
				_ = os.Remove(pathRes)
				lib.NewRes("error copy filesystem", 401, true).Send(w)
				return
			}

			ctx := context.WithValue(r.Context(), "image", imageName)

			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
