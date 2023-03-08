FROM ubuntu:20.04 as builder
# TARGETPLATFORM

ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && apt-get install -y libcurl4-openssl-dev libssl-dev \
    uuid-dev libjsoncpp-dev cmake git build-essential && \
    ln -s /usr/include/jsoncpp/json/ /usr/include/json

WORKDIR /workspace
COPY aliyun-openapi-cpp-sdk ./aliyun-openapi-cpp-sdk
RUN cd aliyun-openapi-cpp-sdk && bash easyinstall.sh core
COPY src ./src
COPY CMakeLists.txt ./CMakeLists.txt
RUN cmake . && make -j

FROM ubuntu:20.04

ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && apt-get install -y libcurl4-openssl-dev libssl-dev \
    uuid-dev libjsoncpp-dev && \
    ln -s /usr/include/jsoncpp/json/ /usr/include/json

WORKDIR /usr/bin/

COPY --from=builder /workspace/aliyun-openapi-cpp-sdk/sdk_build/lib/* /usr/lib/
COPY --from=builder /workspace/bin/cpp-example ./rrsa-example-cpp

CMD /usr/bin/rrsa-example-cpp
