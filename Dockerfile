FROM scratch
ADD build/localgen_amd64_linux_static /localgen
ENTRYPOINT ["/localgen"]