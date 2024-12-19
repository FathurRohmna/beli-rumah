package helper

import (
	"bytes"
	"encoding/base64"
	"html/template"
)

func EncodeBase64URLSafe(input []byte) string {
	return base64.URLEncoding.EncodeToString(input)
}

func RenderTemplate(data map[string]string, htmlTemplate string) (string, error) {
	tmpl, err := template.ParseFiles(htmlTemplate)
	if err != nil {
		return "", err
	}

	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, data); err != nil {
		return "", err
	}

	return rendered.String(), nil
}
