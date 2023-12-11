# My-First-Go-Component

This is a basic example / starter kit to learn about how to build and integrate custom Viam components, using the Viam sensor API. This setup can easily be extended for other component API types listed in the Viam documentation here: [Viam Component APIs](https://docs.viam.com/build/program/apis/#component-apis).


## Useful Go Setup and Build Commands

Some commonly used commands. Maybe helpful maybe not :-)

```go mod init mysensor```

```go mod tidy```


From within the "src" directory run:

```go build -o ../build/mysensor .```

## Add a Local Module

Configure a local module through the web interface or add the following to the RAW JSON.

```
{
  "name": "a-module-name",
  "executable_path": "<-- Path to the sensor binary including the binary -->",
  "type": "local"
}
```



## Configure Component

Add this configuration to the smart machine "components" part either in RAW JSON mode or through the UI and "local component".

```
{
  "name": "gosensor",
  "model": "viam-soleng:sensor:mysensor",
  "type": "sensor",
  "namespace": "rdk",
  "attributes": {
    "setting": 3
  },
  "depends_on": []
}
```