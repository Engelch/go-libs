# go_libs

Go elements often included. This is the successor of the previous debugErrorCE package.

## Crypt

### RSA Key Creation

Keys can be created with the functions:

``
func CreateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error)
``

This function returns a 4096-bit RSA keypair or an error. It contains *syntactic sugar* as the private key always containts the public key. The second functions writes the keys directly to a file.

``
func CreateRSAKeyPair2File(outfileName string) error
``

Keys are automatically stored as a file. By default, the keys are 4096 bit, RSA keys. If the specified filename if `foo`, then the private key is stored in a file `foo` and the public key is stored in a file `foo.pub`. For compatibility reasons, the calls do not try to change the permissions of a file, in particular not of the private key file.

The function returns an error, if either `foo` or `foo.pub` already exists or cannot be created or written to. Both keys are stored in PEM format.

The key creation is compatible with `openssl`. My usual shell functions from the public github.com project (advertisement) [PublicConfigurations](https://github.com/Engelch/PublicConfigurations) allows us to verify if the private and the public key belong together and hereby are also recognised by `openssl`.

```bash
$ tlsRsaPrvFingerprint key1
5d30617645c9cf77d285e4ffd511c407f3aefaa7b2c45ecce6ceba73ccddd27c

$ tlsRsaPubFingerprint key1.pub 
5d30617645c9cf77d285e4ffd511c407f3aefaa7b2c45ecce6ceba73ccddd27c
```
