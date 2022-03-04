# asmkit
Toolkit for genome assembly

## Introduction  

`asmkit` provide some useful tools for genome scaffold, such as [ALLHiC](https://github.com/tangerzhang/ALLHiC).

## Installation
The easiest way to install `asmkit` is to download the latest binary from the [releases](https://github.com/wangyibin/asmkit/releases/latest)  

Build from source with:
```console
go get -u -t -v github.com/wangyibin/asmkit/...
go install github.com/wangyibin/asmkit
```
## Usage
### <kbd>agp2assembly</kbd>

```console
asmkit agp2assembly input.agp output.assembly
```
### <kbd>bam2links</kdb>
```console
asmkit bam2links input.bam output.links
```




## Example
### From bam to juicebox assembly tool
> Depend on [3D-DNA](https://github.com/aidenlab/3d-dna), please first downloand it.

```console
asmkit agp2assembly sample.agp sample.assembly
asmkit bam2links sample.bwa_mem.bam sample.links
bash ~/software/3d-dna/visualize/run-assembly-visualizer.sh sample.assembly sample.links
```
And then import `sample.hic` and `sample.assembly` into [juicebox](https://github.com/aidenlab/Juicebox) to manual curate genome assembly
