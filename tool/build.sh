#!/usr/bin/env bash

src="./unidic-cwj-3.1.1-full"
data_dir=".."
dict="./${data_dir}/uni3.dict"
dest_dir="${data_dir}/dict"
dict_prefix="uni3"

function build_intermediate_files() {
    rm -rf ${dest_dir}
    mkdir -p ${dest_dir}
    split -a 2 -b 100m ${dict} ${dest_dir}/${dict_prefix}
}

go run main.go -dict ${src} -out ${dict}
build_intermediate_files