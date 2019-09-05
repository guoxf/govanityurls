FROM alpine
WORKDIR /workspace
ADD vanity.yaml /workspace
ADD govanityurls /workspace
ENTRYPOINT [ "./govanityurls","-host gxfxx.com" ]