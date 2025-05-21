package helpers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// UploadFileToS3 mengunggah file ke S3 dan mengembalikan URL file yang diunggah
func UploadFileToS3(s3Client *s3.S3, bucketName string, file *multipart.FileHeader, path string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("gagal membuka file: %v", err)
	}
	defer src.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, src)
	if err != nil {
		return "", fmt.Errorf("gagal membaca file: %v", err)
	}

	fileName := fmt.Sprintf("%s/%d_%s", path, time.Now().Unix(), file.Filename)
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(buf.Bytes()),
		ACL:         aws.String("public-read"),
		ContentType: aws.String(file.Header.Get("Content-Type")),
	})
	if err != nil {
		return "", fmt.Errorf("gagal mengunggah file ke S3: %v", err)
	}

	fileURL := fmt.Sprintf("%s/%s/%s", *s3Client.Config.Endpoint, bucketName, fileName)
	return fileURL, nil
}

// DeleteFileFromS3 menghapus file dari S3
func DeleteFileFromS3(s3Client *s3.S3, bucketName, fileURL string) error {
	fileKey := strings.TrimPrefix(fileURL, fmt.Sprintf("%s/%s/", *s3Client.Config.Endpoint, bucketName))
	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileKey),
	})
	if err != nil {
		return fmt.Errorf("gagal menghapus file dari S3: %v", err)
	}
	return nil
}

// validasiFoto memvalidasi file foto yang diunggah
func ValidasiFoto(file *multipart.FileHeader) error {
	if file == nil {
		return fmt.Errorf("foto wajib diisi")
	}

	// Validasi ekstensi file
	ext := strings.ToLower(strings.TrimPrefix(file.Filename, "."))
	if ext != "jpg" && ext != "jpeg" && ext != "png" {
		return fmt.Errorf("format foto tidak valid. Hanya menerima jpg, jpeg, dan png")
	}

	return nil
}
