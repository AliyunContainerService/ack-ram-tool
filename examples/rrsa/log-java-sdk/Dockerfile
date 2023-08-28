FROM maven:3.8.7-openjdk-18-slim as builder
# TARGETPLATFORM

WORKDIR /app
COPY src/main/java/com/alibabacloud/Demo.java ./src/main/java/com/alibabacloud/
COPY pom.xml ./
RUN mvn package


FROM openjdk:18-slim-buster

WORKDIR /app

COPY --from=builder /app/target/demo-1.0-SNAPSHOT-jar-with-dependencies.jar ./

CMD java -jar ./demo-1.0-SNAPSHOT-jar-with-dependencies.jar
