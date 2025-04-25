FROM node:23-slim

RUN useradd -ms /bin/bash student && adduser student sudo
COPY ./dist /home/student/dist

RUN npm install --ignore-scripts -g serve 

WORKDIR /home/student/
USER student

EXPOSE 5137

CMD ["serve", "-s", "dist", "-l", "5137"]