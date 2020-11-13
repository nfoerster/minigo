# Minigo

[Minio](https://min.io/) is a high performance k8s native object storage suite based on the S3-API. Minigo is a small utility application to directly get information from the associated storage suite or to apply file operations on it.

# Configuration

An example configuration file is in the repo:
```yaml
configurations:
    - name: local
      endpoint: localhost:9000
      accesskey: minioadmin
      secretkey: minioadmin
      usessl: false
    ...
``` 

However, exposing secrets is not a good style. Therefore it's also possible to place the secretkey in the environment variable same named as the configuration name. As a second option the configuration file is not needed if the endpoint, accesskey, secretkey and usessl are given directly via arguments.


# Persistent flags

Either config or direct values can be used.

* config - Path to the configuration file.
* configname - Name to specify used configuration name, if not set the first configuration in the configuration file is used.
* endpoint - The endpoint for the minio connection.
* accesskey - The accesskey for minio connection.
* secretkey - The secretkey for minio connection.
* useSSL - Connect with ssl. Defaults to true.

# Usage

## ListBuckets

Prints out all buckets. Example usage:

```console 
minigo listBuckets
    ...
``` 
