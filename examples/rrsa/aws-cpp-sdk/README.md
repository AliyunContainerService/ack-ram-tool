# Usage

Prerequisites:
- You need [aws-cpp-sdk](https://github.com/aws/aws-sdk-cpp) (or at least S3 module of it) installed before using this example.
- cmake
- g++ / clang++

## Build

```bash
update the
cmake ./
make
```

## run
```bash
# NOTE if you want to test it locally:
# set the environment variables according to your ACK Pod
# prepare your Pod token file at /var/run/secrets/ack.alibabacloud.com/rrsa-tokens/token to your local: "$PWD/token"
export ALIBABA_CLOUD_ROLE_ARN=<input>
export ALIBABA_CLOUD_OIDC_TOKEN_FILE="$PWD/token"
export ALIBABA_CLOUD_OIDC_PROVIDER_ARN=<input>

# list objects in bucket
./app <region> <bucket-name>
```
