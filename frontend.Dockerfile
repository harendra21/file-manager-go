FROM ubuntu/apache2:2.4-20.04_beta
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update
RUN apt-get install curl nano -y
RUN apt install -y software-properties-common && add-apt-repository -y ppa:ondrej/php
WORKDIR /var/www/html
ENV NODE_VERSION=16.13.0
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
ENV NVM_DIR=/root/.nvm
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"
RUN node --version
RUN npm --version
COPY ./frontend /var/www/html/
RUN cd /var/www/html && npm i
EXPOSE 80