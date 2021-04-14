# packer-plugin-zstdcompress  
  
packer-plugin-zstdcompress plugin is a copy of the native post-processor 
"compress", which adds support for the ZSTD compression algorithm 

## Installation

In order to install the plugin, simply clone the repo and run the following
commands:

``` bash
mkdir -p ~/.packer.d/plugins
go mod init main
go build -o ~/.packer.d/plugins/
```

## Use

In order to test the plugin, create the following test file
``` hcl 
  post-processors {
    post-processor "artifice" {
      keep_input_artifact = true
      files = [
        "example_archive.tar"
      ]
    }

    post-processor "zstdcompress" {
      keep_input_artifact = false
      format = ".zst"
      output = "example_archive.tar.zst"
    }
  }
```
