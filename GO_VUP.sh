

go version
GO_VER="1.18.1"
ARCH="amd64"
sudo curl -sSL "https://go.dev/dl/go${GO_VER}.linux-${ARCH}.tar.gz" | sudo tar -xz -C /usr/local/
which go
go version

