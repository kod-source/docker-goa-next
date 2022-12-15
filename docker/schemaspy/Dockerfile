# syntax=docker/dockerfile:1.3-labs
FROM openjdk:8u322-slim-bullseye

ARG DRIVER_VERSION="8.0.28"
ARG SCHMASPY_VERSION="6.1.0"

WORKDIR /app

COPY ./schemaspy.properties /app/schemaspy.properties

RUN <<EOF
    apt-get update
    apt-get install -y --no-install-recommends wget=1.21-1+deb11u1
    apt-get install -y --no-install-recommends graphviz=2.42.2-5
    wget -q -O mysql-connector-java.jar "https://repo1.maven.org/maven2/mysql/mysql-connector-java/$DRIVER_VERSION/mysql-connector-java-$DRIVER_VERSION.jar"
    wget -q -O schemaspy.jar "https://github.com/schemaspy/schemaspy/releases/download/v$SCHMASPY_VERSION/schemaspy-$SCHMASPY_VERSION.jar"
    apt-get -y clean
    rm -rf /var/lib/apt/lists/*
EOF

CMD ["java","-jar","schemaspy.jar"]
