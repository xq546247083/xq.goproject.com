package EncrpytTool

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 私钥
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCNJNHiDH21hz6Ru0tA0ybPBzECZIYvkVqzhz+1eBlDBr7kjM1L
EHDelqX1ZX272kG5Xt9Adv6kbcIpRcjiXnYGrt3IhucBlauK9E9V9oxZ7B0n0oUu
VrqyA6YfC9VJM75Hsv1S5Ef7QwL3XajLjJ3EDwz2yP4pu1yWqX4PD0X1wwIDAQAB
AoGAY7qnfmS7El/8ivZfBu/rlQ8MxxcGZWf8QawIDQ5OMzj5+v2uNSLpza1+mDVd
MTXXiMaLUr6B0Oco8Qa5GZILLD8WwQVFeIyV1n/K50dyKKHaqT8bR77JNyd1BYJd
5kV0FFFx2VcEaPBEv+tthWw21DcqwlrHFiXDTqKig4uYUEkCQQCgBVLKwe9IiF6K
/mPzhTTJSQdfpRPxjBKjYRjDpc1o0DAvz7jBxsOUeDLYCsUMGB8MRqowVsYhgZK4
x9aagE1nAkEA4cz/xBSs7DI1xduqydzWHTilzd/LWdaVmSez12wK9SuGGRr0zMzx
2CONh97S3QrrQl311l6exBQN7ia/n31/RQJAKXl3vemJ9UizCF9q1IEf71OoP5fv
lVlysznFS2A73wCmnJ3ACylTI7YLp4cTD1FpKqteDO1QqPqGZrIU4zKB3QJBAJYc
w/hmSmOoKamFFPCoWMwKaegJHNZ32vJ7u4q+cDZ3nem4ywAQS8OGN0QZtZNv++Ee
OB4wv1nZfz1RE1mDhAUCQGX/4PuE0N2US2qFp1nqhP/71dEOzZVN9UJQ9j0k7Uaz
zrBGb8NYHYVvTDO8YYZjzcNvvVF4wtq2iE7Q6pTrVU0=
-----END RSA PRIVATE KEY-----
`)

// 公钥
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCNJNHiDH21hz6Ru0tA0ybPBzEC
ZIYvkVqzhz+1eBlDBr7kjM1LEHDelqX1ZX272kG5Xt9Adv6kbcIpRcjiXnYGrt3I
hucBlauK9E9V9oxZ7B0n0oUuVrqyA6YfC9VJM75Hsv1S5Ef7QwL3XajLjJ3EDwz2
yP4pu1yWqX4PD0X1wwIDAQAB
-----END PUBLIC KEY-----
`)

//RsaEncrypt 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

//RsaDecrypt 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
