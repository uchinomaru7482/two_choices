FROM node:16.3.0

WORKDIR /app
COPY ./two_choices_client /app

RUN cd /app/ && \
    npm install && \
    npm run build && \
    npm run generate
