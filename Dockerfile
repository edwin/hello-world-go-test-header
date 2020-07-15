FROM registry.access.redhat.com/ubi8/go-toolset
ENV GOPATH=$APP_ROOT
ENV GOBIN=$APP_ROOT/bin
ADD . $GOPATH/src/HelloHeader/
RUN go install HelloHeader

FROM registry.access.redhat.com/ubi8/ubi-minimal
COPY --from=0 /opt/app-root/bin/HelloHeader /usr/local/bin/HelloHeader
WORKDIR /usr/local/bin
CMD HelloHeader
EXPOSE 8080