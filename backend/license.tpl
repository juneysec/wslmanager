{
{{ range . }}
    "{{ .Name }}": {
        "LicenseName": "{{ .LicenseName }}",
        "LicenseURL": "{{ .LicenseURL }}"
    },
{{ end }}
}
