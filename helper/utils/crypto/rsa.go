package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/jxncyjq/lib_stardust/common/log"
	"os"
)

// GenerateRSAKeys 生成 RSA 密钥对
func GenerateRSAKeys() error {
	// 生成 RSA 私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// 生成私钥 PEM 文件
	privateKeyFile, err := os.Create("private_key.pem")
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()

	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)
	privateKeyFile.Write(privateKeyPEM)

	// 生成公钥 PEM 文件
	publicKey := &privateKey.PublicKey
	publicKeyFile, err := os.Create("public_key.pem")
	if err != nil {
		return err
	}
	defer publicKeyFile.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicKeyBytes,
		},
	)
	publicKeyFile.Write(publicKeyPEM)

	log.Log.Infof("RSA keys generated and saved to private_key.pem and public_key.pem")
	return nil
}
