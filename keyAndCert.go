package main

import (
	"io/ioutil"
	"log"
)

func createTempFile(name string, content []byte) string {
	tmp, err := ioutil.TempFile("", name)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmp.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmp.Close(); err != nil {
		log.Fatal(err)
	}
	return tmp.Name()
}

const _TLS_CERT = `-----BEGIN CERTIFICATE-----
MIID9zCCAt+gAwIBAgIJANvej/44iwdwMA0GCSqGSIb3DQEBBQUAMFoxCzAJBgNV
BAYTAkNOMRAwDgYDVQQIEwdKaWFuZ3N1MQ8wDQYDVQQHEwZTdXpob3UxDTALBgNV
BAoTBGNteWsxCjAIBgNVBAsTAWExDTALBgNVBAMTBGVyaXMwHhcNMTcwNTAzMTEw
NDQwWhcNMjAwNTAyMTEwNDQwWjBaMQswCQYDVQQGEwJDTjEQMA4GA1UECBMHSmlh
bmdzdTEPMA0GA1UEBxMGU3V6aG91MQ0wCwYDVQQKEwRjbXlrMQowCAYDVQQLEwFh
MQ0wCwYDVQQDEwRlcmlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
teGMP3hKxbBfEvV+u/gQKwJ5cG8bV7twsJv48T18KSjxYpgGitvXh1Y2faFgzYLQ
kxQfZ5owKiegMg5ItH5B4PRKWxfle1pej4epwcKGBibSSf65LFIUitjRAh8bM5sN
S61+PN4Md3uw5GTnaDObcSEJHROF+ls2tQsm+uU/hN68om5sTxLlbhgozpCi7fHP
W1HJea6TGpa5zAuEn0kZ9/e7Yu+tvnhaN2iq73ZZXDiwdKxknuvsBMkAddCM6o7m
jY8o2AlgEfQtmygruiBwVOR7LhawVqyq5WASoVqMq+Zu7qwpRagXEbLn74QRDnsd
/3yn44tov8SoUVuwO6XRQwIDAQABo4G/MIG8MB0GA1UdDgQWBBQgdd4uvSoRlD9W
Vf5FTQhsuhEvATCBjAYDVR0jBIGEMIGBgBQgdd4uvSoRlD9WVf5FTQhsuhEvAaFe
pFwwWjELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0ppYW5nc3UxDzANBgNVBAcTBlN1
emhvdTENMAsGA1UEChMEY215azEKMAgGA1UECxMBYTENMAsGA1UEAxMEZXJpc4IJ
ANvej/44iwdwMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEFBQADggEBAFp0vK9c
f/Sm9GiL3ltlYOcKut4l9GThHdPdjgl5eL2o3ztKk8icDt8SWrmbX3jNbZq2GgnY
LNTLZbOLbf/ZtSxd/mrgke4evhJ+eUKJL0cCgHVXSm1v5Qm0wkuv5/eTXZ0jrKPa
NnYd4mg2mWKt1CUE4JcuYa+Hx2W7YfzUTnYWG7zap6yJU5xSoqRnLYTd9ajGMQ3+
utNp9Rx0hMA7AlEzkEPLOZMVmvc9Xl4ffudOEJbNiH8wmHEvIkIJ/uIFCOzNJkLW
tQBQQnfR8RlDxQ2D5AfBVCg2loMoDTIj1aF+e+eg65K1URy0KYKFTvsfhzz5VfOm
okU94e8LC6JDRU8=
-----END CERTIFICATE-----
`

const _TLS_KEY = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAteGMP3hKxbBfEvV+u/gQKwJ5cG8bV7twsJv48T18KSjxYpgG
itvXh1Y2faFgzYLQkxQfZ5owKiegMg5ItH5B4PRKWxfle1pej4epwcKGBibSSf65
LFIUitjRAh8bM5sNS61+PN4Md3uw5GTnaDObcSEJHROF+ls2tQsm+uU/hN68om5s
TxLlbhgozpCi7fHPW1HJea6TGpa5zAuEn0kZ9/e7Yu+tvnhaN2iq73ZZXDiwdKxk
nuvsBMkAddCM6o7mjY8o2AlgEfQtmygruiBwVOR7LhawVqyq5WASoVqMq+Zu7qwp
RagXEbLn74QRDnsd/3yn44tov8SoUVuwO6XRQwIDAQABAoIBABrHP2/k1RVwFz9i
V6tzJWY/sgOEyEDNfxTxkeBqzgn9VjTO7z+oiH4LinBKv0biuLS+5LTLNcYvGV3l
Tn0MjcSCEySASzAzKPL25V56tmLXemSclUTaPN4IeBVY4Rdi/70hqRFrG+jHlE68
MHneB0fLtlzKkBJBNbfIEZAlCooURaVdgbD4kRVLVhXIkf1NhHd/BBhfwBJVFEtf
xgDFVJ6QvRMLTGhX7SF2aVgjg/NWQacO6H6H8W1tQBe5wbgS3FqDLvN8DaAHGMbp
B07+eRRY6EuT/2WF8mrMvwQCTepKJP7a9xGzwP4SwNNPT+eauPlMsXlw6gBbBP9y
UilsQakCgYEA2GP3NtV1ix34H3a77sTT+Lvz92b5S2iPHYKkR8BVwYVDI2k3EKhy
QP4M/AnSttGMAwUVFUw5If9xhPr+f5nZekzSqy3YwT2dFhHL4V0i27iGGNeiBYqK
QrkaVM3LBr09noduW6NXQu5S6ow3LyGFgiN4sh98K71vhh2z5Oeb/k8CgYEA1yx5
Mn2qxrJiwKTwmBgLI9P93jAAslKDY39yFDy0AUux7G4hX59kzXDEfJsQRywINed/
VmUt15hVWJeBOd2glaG4ZIxr98ldF9rRDHcaZDj7pPHeNYrAlnoH+05JI0X44LhS
a4fDPVAIptCUMVzFn/9qv4GCRBzB0p9IxQwTFM0CgYEApvIpDpMNRrFudsfiIkqU
x3gqtxspBna6w6fBGuJWxyELNUy9gQqUa3QVjHX/rbBpdwL6yx2tjn9Hk7MZmlSW
JvcSKaMLcbsZcKd2Rn7Wn/Hy80OZ6vCBzwyENh03oBFsctxF1klkz2yTaAWZpbEV
EdxQCAymfhB+9FgCLH5MKZ0CgYEAxp93APSfv+rK3aljrgIDxn7ZyU/mVLV2M1Jh
q+yc/NUy2FyQ0gQAuONvfNZmoEw1CA05rjaXqnwjzDOORiwfIAC4hZhZ0dE+7dY/
QpJu+jTfldLum94JCN58n64UNUtSAZ4j9r7Lqr1GPzYqlaZdhuPvVffml5k0EYiX
U81hj8ECgYEAkgNWRtZ92WDUGtnjrGh+rfKB9dMFkoqly9Zm5ooZtIWXqWDmT+xs
10+CAMycPUAAnghLhCLd9RIjttl7DcfF5XQlL8uQf9z1fDwrKq54M7QRXMYWKTM2
Vakj/38HQWJ6jCRhh+Bk74SdWzQCWunughi07BVtAwLekYOskQDzlEc=
-----END RSA PRIVATE KEY-----
`
