FROM node:23-slim

WORKDIR /app
COPY . .

RUN npm install \
    && npm install -g serve\
    && npm run build

EXPOSE 5137

CMD ["serve", "-s", "dist", "-l", "5137"]