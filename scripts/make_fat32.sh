#!/bin/bash

IMAGE=fat32.img

# note: a fat32 image needs to be >= 34 MiB
SIZE=64

rm -f $IMAGE

# create an image of zeroes
dd if=/dev/zero of=$IMAGE bs=1M count=$SIZE

# format it as a fat32 file system
mkfs.vfat -n "TEC-1G" -F 32 $IMAGE

# note: uses mtools
# sudo apt-get install mtools
mcopy -i $IMAGE ./content/* ::/
mdir -i $IMAGE ::
