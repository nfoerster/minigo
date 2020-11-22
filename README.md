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

```
--config - Path to the configuration file.
--configname - Name to specify used configuration name, if not set the first configuration in the configuration file is used.
--endpoint - The endpoint for the minio connection.
--accesskey - The accesskey for minio connection.
--secretkey - The secretkey for minio connection.
--useSSL - Connect with ssl. Defaults to true.
```

# Usage

## ListBuckets

Prints out all buckets. Example usage:
```console 
minigo listBuckets
``` 

Optional arguments:
```
-ts - Show creation timestamps of buckets
```

## CreateBucket

Creates a given bucket by name. Usage:
```console 
minigo createBucket <bucketname>
``` 

Example usage to create a bucket named test:
```console 
minigo createBucket test
``` 

Optional arguments:
```
-location - Specify the location of the bucket to create
```

## BucketExists

Just prints out if the given bucket exists. Usage:
```console 
minigo bucketExists <bucketname>
``` 

Example usage to check if test bucket exists:
```console 
minigo bucketExists test
``` 

## RemoveBucket

Removes the bucket if exists. Usage:
```console 
minigo removeBucket <bucketname>
``` 

Example usage to remove the bucket test:
```console 
minigo removeBucket test
``` 

## ListObjects

List all objects in a bucket or given subfolder. Usage:
```console 
minigo listObjects <bucketname>
``` 

Optional arguments:
```
-prefix - Specify the subpath in the bucket to list
-recursive - Recursively searches all folders and subfolders 
-ts - Additionally print out latest modification timestamps
```

Example usage to list all objects recursively in bucket test within folder3:
```console 
minigo listObjects test -prefix folder3 -recursive
``` 

## RemoveObjects

Removes all objects in a bucket or given subfolder. Usage:
```console 
minigo removeObjects <bucketname> <objectname>
``` 

Remove all objects in folder3 in bucket test:
```console 
minigo removeObjects test folder3
``` 

Removes web.go inside folder3 in bucket test:
```console 
minigo removeObjects test folder3/web.go
``` 

## GetObjects

Downloads the object or all objects to the local named structure given in destination. Usage:
```console 
minigo getObjects <bucketname> <objectname> <destination>
``` 

Downloads all objects in folder3 in bucket test to minigodl folder locally:
```console 
minigo getObjects test folder3 minigodl
``` 

Downloads web.go in folder3 in bucket test to minigodl/web.go locally:
```console 
minigo getObjects test folder3/web.go minigodl/
``` 

Downloads web.go in folder3 in bucket test to gui.go locally:
```console 
minigo getObjects test folder3/web.go gui.go
``` 

## CopyObjects

Copy the object or all objects to another location. Usage:
```console 
minigo copyObjects <bucketname> <objectname> <destination>
``` 

Optional arguments:
```
-bucket - Specifies another bucket as the destination bucket
```

Copies all objects in folder3 in bucket test to minigodl folder in bucket test:
```console 
minigo copyObjects test folder3 minigodl
``` 

Copies web.go in folder3 in bucket test to minigodl/web.go in bucket test:
```console 
minigo copyObjects test folder3/web.go minigodl/
``` 

Copies web.go in folder3 in bucket test to gui.go in bucket test:
```console 
minigo copyObjects test folder3/web.go gui.go
``` 

Copies web.go in folder3 in bucket test to gui.go in bucket anothertest:
```console 
minigo copyObjects test folder3/web.go gui.go -bucket anothertest
``` 

Copies all objects from bucket test to bucket anothertest:
```console 
minigo copyObjects test . . -bucket anothertest
``` 

Copies all objects from bucket test to folder anotherFolder in bucket anothertest:
```console 
minigo copyObjects test . anotherFolder -bucket anothertest
``` 

## MirrorBucket

Mirrors a bucket and all of its content to another bucket. If this bucket does not exist, create it. Usage:
```console 
minigo mirrorBucket <sourceBucketName> <targetBucketName>
``` 

Copies all objects bucket test to the new bucket anothertest:
```console 
minigo mirrorBucket test anothertest
``` 

Optional arguments:
```
-location - Specify the location of the target bucket
```
