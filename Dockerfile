FROM maven:3.8.4-openjdk-17 as builder
# Building image
WORKDIR /app
# Forwarding everything except code

COPY . /app/.

# COPY pom.xml /app/.

# Collecting Deps
RUN --mount=type=cache,target=/root/.m2 mvn clean package  -Dmaven.test.skip


FROM eclipse-temurin:17-jre-alpine
WORKDIR /app
COPY --from=builder /app/target/*.jar /app/*.jar
EXPOSE 8181
ENTRYPOINT ["java", "-jar", "/app/*.jar"]