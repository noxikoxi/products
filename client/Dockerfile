FROM node:23-slim AS builder

WORKDIR /app

ENV VITE_API_URL="http://localhost:1323"

COPY package*.json ./
COPY vite.config.ts ./
COPY tsconfig.json ./
COPY tsconfig.app.json ./
COPY tsconfig.node.json ./

RUN npm install --ignore-scripts

COPY public ./public
COPY src ./src
COPY index.html ./

RUN npm run build

FROM nginx:alpine AS runner

COPY --from=builder /app/dist /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf

RUN mkdir -p /run/nginx /var/cache/nginx/client_temp && \
    chown -R nginx:nginx /var/cache/nginx /usr/share/nginx/html /run

COPY entrypoint.sh /docker-entrypoint.d/
RUN chmod +x /docker-entrypoint.d/entrypoint.sh

USER nginx

EXPOSE 80

CMD ["nginx","-g","daemon off;"]