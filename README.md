# Moby Container Manager

Sebuah aplikasi Go sederhana yang menggunakan komponen Moby/Docker untuk mengelola container. Project ini mendemonstrasikan penggunaan praktis dari Moby Project dalam aplikasi Go.

## Fitur

- List semua container (running dan stopped)
- Start container
- Stop container
- Inspect detail container
- Remove container
- Pull Docker images
- Interface CLI yang user-friendly

## Instalasi

### Prerequisites
- Go 1.19 atau lebih baru
- Docker daemon yang berjalan

### Build dari Source
```bash
git clone <repository-url>
cd moby-container-manager
go mod download
go build -o mcm cmd/main.go
```

## Penggunaan

### List Container
```bash
./mcm list
./mcm list --all  # termasuk stopped containers
```

### Start Container
```bash
./mcm start <container-id-or-name>
```

### Stop Container
```bash
./mcm stop <container-id-or-name>
```

### Inspect Container
```bash
./mcm inspect <container-id-or-name>
```

### Remove Container
```bash
./mcm remove <container-id-or-name>
```

### Pull Image
```bash
./mcm pull <image-name>
```

## Contoh Penggunaan

```bash
# List semua running containers
./mcm list

# List semua containers (termasuk yang stopped)
./mcm list --all

# Start sebuah container
./mcm start my-container

# Stop sebuah container
./mcm stop my-container

# Inspect detail container
./mcm inspect my-container

# Pull image baru
./mcm pull nginx:latest
```

## Struktur Project

```
moby-container-manager/
├── cmd/
│   └── main.go           # Entry point aplikasi
├── internal/
│   ├── client/           # Docker client wrapper
│   ├── commands/         # Command implementations
│   └── models/           # Data models
├── go.mod
├── go.sum
└── README.md
```

## Dependensi

Project ini menggunakan:
- `github.com/docker/docker` - Docker client library (bagian dari Moby)
- `github.com/spf13/cobra` - CLI framework
- `github.com/spf13/viper` - Configuration management

## Kontribusi

1. Fork repository
2. Buat feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push ke branch (`git push origin feature/amazing-feature`)
5. Buat Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.