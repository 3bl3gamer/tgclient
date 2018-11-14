package mtproto

import (
	"crypto/aes"
	"crypto/rsa"
	sha1lib "crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"math/big"
	"math/rand"
	"time"

	"github.com/ansel1/merry"
	"golang.org/x/crypto/pbkdf2"
)

const (
	telegramPublicKey_N  = "24403446649145068056824081744112065346446136066297307473868293895086332508101251964919587745984311372853053253457835208829824428441874946556659953519213382748319518214765985662663680818277989736779506318868003755216402538945900388706898101286548187286716959100102939636333452457308619454821845196109544157601096359148241435922125602449263164512290854366930013825808102403072317738266383237191313714482187326643144603633877219028262697593882410403273959074350849923041765639673335775605842311578109726403165298875058941765362622936097839775380070572921007586266115476975819175319995527916042178582540628652481530373407"
	telegramPublicKey_E  = 65537
	telegramPublicKey_FP = -4344800451088585951 // 14101943622620965665 as int64
)

var telegramPublicKey rsa.PublicKey

func init() {
	telegramPublicKey.N, _ = new(big.Int).SetString(telegramPublicKey_N, 10)
	telegramPublicKey.E = telegramPublicKey_E
}

func sha1(data []byte) []byte {
	r := sha1lib.Sum(data)
	return r[:]
}

func sha256some(buffers ...[]byte) []byte {
	s := sha256.New()
	for _, buf := range buffers {
		s.Write(buf)
	}
	return s.Sum(nil)
}

func xor(dst, src []byte) {
	for i := range dst {
		dst[i] = dst[i] ^ src[i]
	}
}

func bigIntPaddedBytes(num *big.Int, size int) []byte {
	buf := make([]byte, size)
	numBuf := num.Bytes()
	copy(buf[size-len(numBuf):], numBuf)
	return buf
}

func doRSAencrypt(em []byte) []byte {
	z := make([]byte, 255)
	copy(z, em)

	c := new(big.Int)
	c.Exp(new(big.Int).SetBytes(z), big.NewInt(int64(telegramPublicKey.E)), telegramPublicKey.N)

	res := make([]byte, 256)
	copy(res, c.Bytes())

	return res
}

func splitPQ(pq *big.Int) (p1, p2 *big.Int) {
	value_0 := big.NewInt(0)
	value_1 := big.NewInt(1)
	value_15 := big.NewInt(15)
	value_17 := big.NewInt(17)
	rndmax := big.NewInt(0).SetBit(big.NewInt(0), 64, 1)

	what := big.NewInt(0).Set(pq)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	g := big.NewInt(0)
	i := 0
	for !(g.Cmp(value_1) == 1 && g.Cmp(what) == -1) {
		q := big.NewInt(0).Rand(rnd, rndmax)
		q = q.And(q, value_15)
		q = q.Add(q, value_17)
		q = q.Mod(q, what)

		x := big.NewInt(0).Rand(rnd, rndmax)
		whatnext := big.NewInt(0).Sub(what, value_1)
		x = x.Mod(x, whatnext)
		x = x.Add(x, value_1)

		y := big.NewInt(0).Set(x)
		lim := 1 << (uint(i) + 18)
		j := 1
		flag := true

		for j < lim && flag {
			a := big.NewInt(0).Set(x)
			b := big.NewInt(0).Set(x)
			c := big.NewInt(0).Set(q)

			for b.Cmp(value_0) == 1 {
				b2 := big.NewInt(0)
				if b2.And(b, value_1).Cmp(value_0) == 1 {
					c.Add(c, a)
					if c.Cmp(what) >= 0 {
						c.Sub(c, what)
					}
				}
				a.Add(a, a)
				if a.Cmp(what) >= 0 {
					a.Sub(a, what)
				}
				b.Rsh(b, 1)
			}
			x.Set(c)

			z := big.NewInt(0)
			if x.Cmp(y) == -1 {
				z.Add(what, x)
				z.Sub(z, y)
			} else {
				z.Sub(x, y)
			}
			g.GCD(nil, nil, z, what)

			if (j & (j - 1)) == 0 {
				y.Set(x)
			}
			j = j + 1

			if g.Cmp(value_1) != 0 {
				flag = false
			}
		}
		i = i + 1
	}

	p1 = big.NewInt(0).Set(g)
	p2 = big.NewInt(0).Div(what, g)

	if p1.Cmp(p2) == 1 {
		p1, p2 = p2, p1
	}

	return
}

func makeGAB(g int32, g_a, dh_prime *big.Int) (b, g_b, g_ab *big.Int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndmax := big.NewInt(0).SetBit(big.NewInt(0), 2048, 1)
	b = big.NewInt(0).Rand(rnd, rndmax)
	g_b = big.NewInt(0).Exp(big.NewInt(int64(g)), b, dh_prime)
	g_ab = big.NewInt(0).Exp(g_a, b, dh_prime)

	return
}

func generateAES(msg_key, auth_key []byte, decode bool) ([]byte, []byte) {
	var x int
	if decode {
		x = 8
	} else {
		x = 0
	}
	aes_key := make([]byte, 0, 32)
	aes_iv := make([]byte, 0, 32)
	t_a := make([]byte, 0, 48)
	t_b := make([]byte, 0, 48)
	t_c := make([]byte, 0, 48)
	t_d := make([]byte, 0, 48)

	t_a = append(t_a, msg_key...)
	t_a = append(t_a, auth_key[x:x+32]...)

	t_b = append(t_b, auth_key[32+x:32+x+16]...)
	t_b = append(t_b, msg_key...)
	t_b = append(t_b, auth_key[48+x:48+x+16]...)

	t_c = append(t_c, auth_key[64+x:64+x+32]...)
	t_c = append(t_c, msg_key...)

	t_d = append(t_d, msg_key...)
	t_d = append(t_d, auth_key[96+x:96+x+32]...)

	sha1_a := sha1(t_a)
	sha1_b := sha1(t_b)
	sha1_c := sha1(t_c)
	sha1_d := sha1(t_d)

	aes_key = append(aes_key, sha1_a[0:8]...)
	aes_key = append(aes_key, sha1_b[8:8+12]...)
	aes_key = append(aes_key, sha1_c[4:4+12]...)

	aes_iv = append(aes_iv, sha1_a[8:8+12]...)
	aes_iv = append(aes_iv, sha1_b[0:8]...)
	aes_iv = append(aes_iv, sha1_c[16:16+4]...)
	aes_iv = append(aes_iv, sha1_d[0:8]...)

	return aes_key, aes_iv
}

func doAES256IGEencrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, merry.Errorf("AES256IGE: data too small to encrypt: %d < %d", len(data), aes.BlockSize)
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, merry.Errorf("AES256IGE: data not divisible by block size: %d %% %d != 0", len(data), aes.BlockSize)
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, iv[:aes.BlockSize])
	copy(y, iv[aes.BlockSize:])
	encrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(x, data[i:i+aes.BlockSize])
		block.Encrypt(t, x)
		xor(t, y)
		x, y = t, data[i:i+aes.BlockSize]
		copy(encrypted[i:], t)
		i += aes.BlockSize
	}

	return encrypted, nil
}

func doAES256IGEdecrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, merry.Errorf("AES256IGE: data too small to decrypt: %d < %d", len(data), aes.BlockSize)
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, merry.Errorf("AES256IGE: data not divisible by block size: %d %% %d != 0", len(data), aes.BlockSize)
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, iv[:aes.BlockSize])
	copy(y, iv[aes.BlockSize:])
	decrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(y, data[i:i+aes.BlockSize])
		block.Decrypt(t, y)
		xor(t, x)
		y, x = t, data[i:i+aes.BlockSize]
		copy(decrypted[i:], t)
		i += aes.BlockSize
	}

	return decrypted, nil

}

func calcInputCheckPasswordSRP(
	algo TL_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow,
	accPassword TL_account_password,
	password string,
	randFunc func([]byte) (int, error),
	logDebug func(string, ...interface{}),
) (TL, error) {
	logDebug(" --- SRP calculation start --- ")
	logDebug("algo.Salt1 (client): %#v", algo.Salt1)
	logDebug("algo.Salt2 (server): %#v", algo.Salt2)
	logDebug("algo.G:              %#v", algo.G)
	logDebug("algo.P:              %#v", algo.P)
	logDebug("accPassword.SrpB:    %#v", accPassword.SrpB)
	logDebug("accPassword.SrpID:   %#v", accPassword.SrpID)
	logDebug("password:            %#v", password)
	defer logDebug(" --- SRP calculation end --- ")

	if len(password) == 0 {
		return nil, merry.New("password is empty")
	}
	// TODO: check g and p
	// https://github.com/tdlib/td/blob/d9a18a064fa99130dc9214fb6131ba59e5660892/td/telegram/PasswordManager.cpp#L79
	// DhHandshake::check_config(g, p, ...

	clientSalt := algo.Salt1
	serverSalt := algo.Salt2
	gVal := algo.G
	pBuf := algo.P
	BBuf := accPassword.SrpB
	srpID := accPassword.SrpID

	if len(BBuf) != 256 {
		return nil, merry.Errorf("wrong SrpB size, expected 256 bytes, got %d", len(BBuf))
	}

	gNum := new(big.Int).SetInt64(int64(gVal))
	gBuf := bigIntPaddedBytes(gNum, 256)
	pNum := new(big.Int).SetBytes([]byte(pBuf))
	BNum := new(big.Int).SetBytes([]byte(BBuf))

	if BNum.Cmp(pNum) != -1 {
		return nil, merry.Errorf("expected SrpB < P, got: SrpB = %s, P = %s", BNum, pNum)
	}

	// calc_password_hash
	buf := sha256some(clientSalt, []byte(password), clientSalt)
	buf = sha256some(serverSalt, buf, serverSalt)
	hash := pbkdf2.Key(buf, clientSalt, 100000, 64, sha512.New)
	xBuf := sha256some(serverSalt, hash, serverSalt)
	xNum := new(big.Int).SetBytes(xBuf)
	// /calc_password_hash

	aBuf := make([]byte, 2048/8)
	_, err := randFunc(aBuf)
	if err != nil {
		return nil, merry.Wrap(err)
	}
	aNum := new(big.Int).SetBytes(aBuf)
	logDebug("a: %s %#v", aNum, aBuf)

	ANum := new(big.Int).Exp(gNum, aNum, pNum)
	ABuf := bigIntPaddedBytes(ANum, 256)

	uBuf := sha256some(ABuf, BBuf)
	uNum := new(big.Int).SetBytes(uBuf)
	kBuf := sha256some(pBuf, gBuf)
	kNum := new(big.Int).SetBytes(kBuf)

	vNum := new(big.Int).Exp(gNum, xNum, pNum)
	kvNum := new(big.Int).Mul(kNum, vNum)
	kvNum.Mod(kvNum, pNum)
	tNum := new(big.Int).Sub(BNum, kvNum)
	if tNum.Sign() == -1 {
		logDebug("LZ!!! %s", tNum)
		tNum.Add(tNum, pNum)
	}
	expNum := new(big.Int).Mul(uNum, xNum)
	expNum.Add(expNum, aNum)

	SNum := new(big.Int).Exp(tNum, expNum, pNum)
	SBuf := bigIntPaddedBytes(SNum, 256)
	KBuf := sha256some(SBuf)

	h1 := sha256some(pBuf)
	h2 := sha256some(gBuf)
	xor(h1, h2)
	clientSaltHash := sha256some(clientSalt)
	serverSaltHash := sha256some(serverSalt)
	MBuf := sha256some(h1, clientSaltHash, serverSaltHash, ABuf, BBuf, KBuf)
	logDebug("srpID: %#v", srpID)
	logDebug("ABuf:  %#v", ABuf)
	logDebug("MBuf:  %#v", MBuf)

	return TL_inputCheckPasswordSRP{SrpID: srpID, A: ABuf, M1: MBuf}, nil
}
