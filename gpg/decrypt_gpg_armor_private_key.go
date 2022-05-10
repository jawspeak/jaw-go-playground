package main

// Initially from https://gist.github.com/jyap808/8310117
// Slightly modified for updated x/crypto package imports

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

func main() {
	// Read armored private key into type EntityList
	// An EntityList contains one or more Entities.

	entitylist, err := openpgp.ReadArmoredKeyRing(bytes.NewBufferString(publicKey))
	if err != nil {
		log.Fatal(err)
	}

	var notExpiredEntities openpgp.EntityList

	currentTime := time.Now()
	fmt.Println("The len of the EntityList", len(entitylist))
	for i, entity := range entitylist {
		fmt.Printf("i=%d has %d length of identities\n", i, len(entity.Identities))
		fmt.Println("entity.PrimaryKey.CreationTime", entity.PrimaryKey.CreationTime)
		if entity.PrivateKey != nil {
			fmt.Println("entity.PrivateKey.CreationTime", entity.PrivateKey.CreationTime)
		} else {
			fmt.Println("entity.PrivateKey = nil, this is only a public key")
		}

		foundExpired := false

		// There will only be 1 identity on our key but this is here just in case
		for _, identity := range entity.Identities {
			fmt.Printf("Identity name: \"%s\" has %d len Signatures. Is key expired? %v\n", identity.Name, len(identity.Signatures), identity.SelfSignature.KeyExpired(currentTime))
			fmt.Println("identity.SelfSignature.CreationTime", identity.SelfSignature.CreationTime)
			for k, sig := range identity.Signatures {
				fmt.Printf("\t\tk=%d created at %v. Is sig key expired? %v\n", k, sig.CreationTime, sig.KeyExpired(currentTime))
				fmt.Println("sig.CreationTime", sig.CreationTime)
				if sig.KeyExpired(currentTime) {
					// technically it may not matter if another of these are expired, not sure - but our key from keywhiz doesn't have these other signatures so it's ok either way.
					foundExpired = true
				}
			}

			// There is a bug report that this should look not at the signature expiration but
			// the key expiration. But for our use the unpatched library is ok.
			// https://github.com/golang/go/issues/22312
			if identity.SelfSignature.KeyExpired(currentTime) {
				foundExpired = true
			}
		}
		if !foundExpired {
			notExpiredEntities = append(notExpiredEntities, entity)
		}
	}
	fmt.Println()

	fmt.Println("Not expired entities len", len(notExpiredEntities))

	// example of decrypting

	// Decrypt armor encrypted message using private key
	decbuf := bytes.NewBuffer([]byte(encryptedMessage))
	result, err := armor.Decode(decbuf)
	if err != nil {
		log.Fatal(err)
	}

	md, err := openpgp.ReadMessage(result.Body, entitylist, nil, nil)
	if err != nil {
		fmt.Println("error reading message", err)
	}

	bytes, err := ioutil.ReadAll(md.UnverifiedBody)
	fmt.Println("md:", string(bytes))

}

// BELOW IS WRONG
// sec   1024R/1D489768 2014-01-08
// uid                  Golang Test (Private key has no password) <golangtest@test.com>
// ssb   1024R/A46346C7 2014-01-08

const publicKey = `-----BEGIN PGP PRIVATE KEY BLOCK-----

xsBNBGF5tikBCADEh4ozwqOvk2eUnviqj8/ma9MZwqG/hw3Pn/Is2+ktid6p8qwp
LolK2ee3na7086Vb/DH108ig6bLgL14UXrzVOEvdBNIUcGuBQEScO+tbDYGmKYtK
pDFBhilAxqOhqXLJ+MMIkok+3XXs61bE1La9JNZsxTouZMeQhkXX+Tq7ALOxpGtR
WkdmA1Oh5NsMxZr/n/MBPVYQyix/qZHHNs39jUDJkI1NePY6/kW966gX13S3lcJ+
qhvJJ7xY4hvi+oeue1B089sGMQ98NdLyGBNsGR+KyJEczYj43+LIOVKonSegmywq
+ITPf5+/YEvOU7xCL1SxczdyzCG5SqxP/KurABEBAAHNnXRlbGxlciBwcm9kdWN0
aW9uIHNmcy1hY2gtc2lnbmluZy0yMDIxLTEwLTI3IGtleSAoU2lnbmluZyBhbmQg
ZW5jcnlwdGlvbiBrZXkgZm9yIHNmcy1hY2gtc2lnbmluZy0yMDIxLTEwLTI3IGlu
dGVncmF0aW9uKSA8dGVsbGVyLXByb2R1Y3Rpb25AZ3BnLnNxdWFyZXVwLmNvbT7C
wHoEEwEIAC4FAmF5tikJEKACXynrwtIMAhsDBQkDwmcAAhkBBgsJCAcDAgYVCAIJ
CgsDFgIBAABR+wgAN1heHzvYjTr2ce32V1SCjqyZI8irjNnB5lp9E/5JKan7sWvr
j3QPhh4IDglFpuBh7mKESK376NIfzMcALLpqwCetRkNpqov4einkPnz47yLfrT2H
ynFyzq9cQptRfbhyr6J88LI793PuVgjOFGegiNdo79ROoGVYzxEsZ7Y2E+vB/Wru
Qvh0OXLNp05G1kkvPpPs8VzHpNmIKTEP0CFccr3DdfJQA1uj81OsacBxVtPFaD/0
iELaSTjgxomxHnP2QA0OFWfYCU9KuUQ7kpbjSVZgZs1P53KymeNIuhIWWJe+c+9x
KWTR0vYwAmZjmn9paPrrhKNwiOqQFNlzsEi0o87ATQRhebYpAQgA0dSueHJL+nGH
EpmUhmkEkizpF7+1eSZCeNytwIQS0l/w7TsCIOQlsi4qk6TAH/cx1Iz69LLbQEj5
H/eLbq5JLj8SuzDh7q4bXH1OuV6rp9KBS7m73WlW8SenMClcaqHVmt9/AkF/KPh9
Vgeh+p1wH5nWD6AcZswft/IbdyWTYRatglb4nQKGxjwgONtJcfI+IssBYHIW+XEk
7/A2aBlWaASVyPIAjFDvqzxRLRHAApZG2o2Rd5wux49Vu7AvB11Ek0al0kvJ4glO
LFrs+iObLimIcfpN9wv0Uwl5dkj+6sz7J8xR5/6WlY0lq0HMz78iV7eEGzkPlSH4
+SrC+I3vaQARAQABwsBlBBgBCAAZBQJhebYpCRCgAl8p68LSDAIbDAUJA8JnAAAA
PF0IAJFne3RKmRvIE8oH7TXu6Rj+vfZwxvfCojR0MxsjfMIZDi5LFUV1dBLEc65F
QefV1KK7/xADAug2H+1afRqWwS0lzSeqKnscFsfNbMY6FCoKsUuxo1Avxh1WqYmS
9nhGGw39YRPZww8nTgD6PAxw7SHFzkxpfxyVMXYh8Ksb1WFYsniWvChiTzMRa5tv
PG1obiMUL2t8LA67PID2Rxf6PWRqBtBEi+zvKBaCT7aXc+KdnwwS4YjMAu7+78o+
IVhk4z3ejr1apbiVhe1eVf612LQ6Aus690wW5RORqto5gs0MxcteQfJUMZ9+nuHF
VlJ5GrwZhRVI8uiCR9wBFQygzlc=
=H0PV
-----END PGP PRIVATE KEY BLOCK-----
`

// Encrypted by public key message for
// 1024R/A46346C7 2014-01-08 "Golang Test (Private key has no password) <golangtest@test.com>"
const encryptedMessage = `-----BEGIN PGP MESSAGE-----                  
Version: GnuPG v1

hIwDa78K16RjRscBA/9rEXnQ06mAOwrhTXwkps0cYhOBc/dX7BtdEPoRLBd0zNpr
hlU2YPcmQjBaL2Zi2E64K9Ud+R3D4RHt78a1145jjPFBMpWg1z2SMgERaRtwT2++
AinkXavDTuwoyPD7X3z4Jyp+aPcvReUQhq7idb2Kl5cjnXe1Z31a9wy4wMHEItJh
ARDEtteoYHsszGqTdIQOuGrI5P/V4biteh5JpxWqeTWDiY9HR2O31kwRhh9bV3J/
0kFgWpTRqfaSmu8ZiHtfQzTGCHpbLhz2IQgOSjBevzZACbn0z7h5Ro+EfAsonGfH
3w==
=xUUr
-----END PGP MESSAGE-----`
