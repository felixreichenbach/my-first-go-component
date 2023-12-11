# My-First-Go-Component

This is a basic example / starter kit template helping you to learn and get started easily with integrating your particular hardware into the Viam ecosystem. This starter kit uses a sensor component to keep it simple but can also easily be used as foundation for any other component. All you have to do is to replace the sensor specific "Readings" method with the methods required by the component you intend to integrate. 

You can find further component APIs in the Viam documentation here: [Viam Component APIs](https://docs.viam.com/build/program/apis/#component-apis).

## Useful Go Setup and Build Commands

Some commonly used commands. Maybe helpful maybe not :-)

```go mod init mysensor```

```go mod tidy```


From within the "src" directory run:

```go build -o ../bin/mysensor .```

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
      "name": "mysensor",
      "model": "viam-soleng:go-resources:gosensor",
      "type": "sensor",
      "namespace": "rdk",
      "attributes": {
        "setting": 2
      },
      "depends_on": []
    }
```
